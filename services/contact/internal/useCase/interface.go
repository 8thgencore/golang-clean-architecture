package useCase

import (
	"app/pkg/type/context"
	"app/pkg/type/queryParameter"
	"app/services/contact/internal/domain/contact"
	"app/services/contact/internal/domain/group"

	"github.com/google/uuid"
)

type Contact interface {
	Create(c context.Context, constacts ...*contact.Contact) ([]*contact.Contact, error)
	Update(c context.Context, contact contact.Contact) (*contact.Contact, error)
	Delete(c context.Context, ID uuid.UUID /*Тут можно передавать фильтр*/) error

	ContactReader
}

type ContactReader interface {
	List(c context.Context, parameter queryParameter.QueryParameter) ([]*contact.Contact, error)
	ReadByID(c context.Context, ID uuid.UUID) (response *contact.Contact, err error)
	Count(c context.Context /*Тут можно передавать фильтр*/) (uint64, error)
}

type Group interface {
	Create(c context.Context, creator group.Creator) (group.Reader, error)
	Update(c context.Context, updater group.Updater) (group.Reader, error)
	Delete(c context.Context, deleter group.Deleter /*Тут можно передавать фильтр*/) error

	GroupReader
	ContactInGroup
}

type GroupReader interface {
	List(c context.Context, paramete queryParameter.QueryParameter) ([]group.Reader, error)
	ReadByID(c context.Context, ID uuid.UUID) (group.Reader, error)
	Count(c context.Context /*Тут можно передавать фильтр*/) (uint64, error)
}

type ContactInGroup interface {
	CreateContactIntoGroup(c context.Context, groupID uuid.UUID, contacts ...*contact.Contact) ([]*contact.Contact, error)
	AddContactToGroup(c context.Context, groupID, contactID uuid.UUID) error
	DeleteContactFromGroup(c context.Context, groupID, contactID uuid.UUID) error
}
