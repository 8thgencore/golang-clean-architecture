package group

import (
	"app/services/contact/internal/domain/group/description"
	"app/services/contact/internal/domain/group/name"
	"time"

	"github.com/google/uuid"
)

type group struct {
	id           uuid.UUID
	createdAt    time.Time
	modifiedAt   time.Time
	name         name.Name
	description  description.Description
	contactCount uint64
}

func (g group) ContactCount() uint64 {
	return g.contactCount
}

func (g group) ID() uuid.UUID {
	return g.id
}

func (g group) CreatedAt() time.Time {
	return g.createdAt
}

func (g group) ModifiedAt() time.Time {
	return g.modifiedAt
}

func (g group) Name() name.Name {
	return g.name
}

func (g group) Description() description.Description {
	return g.description
}
