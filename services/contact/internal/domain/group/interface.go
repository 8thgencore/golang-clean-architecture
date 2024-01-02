package group

import (
	"app/services/contact/internal/domain/group/description"
	"app/services/contact/internal/domain/group/name"
	"time"

	"github.com/google/uuid"
)

type Creator interface {
	ID() uuid.UUID
	CreatedAt() time.Time
	ModifiedAt() time.Time
	Name() name.Name
	Description() description.Description
}

type Reader interface {
	ID() uuid.UUID
	CreatedAt() time.Time
	ModifiedAt() time.Time
	Name() name.Name
	Description() description.Description
	ContactCount() uint64
}

type Updater interface {
	ID() uuid.UUID
	ModifiedAt() time.Time
	Name() name.Name
	Description() description.Description
}

type Deleter interface {
	ID() uuid.UUID
	ModifiedAt() time.Time
}
