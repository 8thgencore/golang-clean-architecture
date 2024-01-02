package storage

import (
	"app/pkg/type/context"
	"app/pkg/type/queryParameter"
	"app/services/contact/internal/domain/contact"
	"app/services/contact/internal/domain/group"

	"github.com/google/uuid"
)

type Storage interface {
	Contact
	Group
}

type Contact interface {
	CreateContact(ctx context.Context, contacts ...*contact.Contact) ([]*contact.Contact, error)
	UpdateContact(ctx context.Context, ID uuid.UUID, updateFn func(c *contact.Contact) (*contact.Contact, error)) (*contact.Contact, error)
	DeleteContact(ctx context.Context, ID uuid.UUID) error

	ContactReader
}

type ContactReader interface {
	ListContact(ctx context.Context, parameter queryParameter.QueryParameter) ([]*contact.Contact, error)
	ReadContactByID(ctx context.Context, ID uuid.UUID) (response *contact.Contact, err error)
	CountContact(ctx context.Context /*Тут можно передавать фильтр*/) (uint64, error)
}

type Group interface {
	CreateGroup(ctx context.Context, creator group.Creator) (group.Reader, error)
	UpdateGroup(ctx context.Context, updater group.Updater) (group.Reader, error)
	DeleteGroup(ctx context.Context, deleter group.Deleter /*Тут можно передавать фильтр*/) error

	GroupReader
	ContactInGroup
}

type GroupReader interface {
	ListGroup(ctx context.Context, parameter queryParameter.QueryParameter) ([]group.Reader, error)
	ReadGroupByID(ctx context.Context, ID uuid.UUID) (group.Reader, error)
	CountGroup(ctx context.Context /*Тут можно передавать фильтр*/) (uint64, error)
}

type ContactInGroup interface {
	CreateContactIntoGroup(ctx context.Context, groupID uuid.UUID, in ...*contact.Contact) ([]*contact.Contact, error)
	DeleteContactFromGroup(ctx context.Context, groupID, contactID uuid.UUID) error
	AddContactsToGroup(ctx context.Context, groupID uuid.UUID, contactIDs ...uuid.UUID) error
}
