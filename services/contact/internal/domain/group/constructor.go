package group

import (
	"app/services/contact/internal/domain/group/description"
	"app/services/contact/internal/domain/group/name"
	"time"

	"github.com/google/uuid"
)

func NewCreator(name name.Name, description description.Description) Creator {
	var timeNow = time.Now().UTC()
	return &group{
		id:          uuid.New(),
		name:        name,
		description: description,
		createdAt:   timeNow,
		modifiedAt:  timeNow,
	}
}

func NewReader(id uuid.UUID, createdAt time.Time, modifiedAt time.Time, name name.Name, description description.Description, contactCount uint64) Reader {
	return &group{
		id:           id,
		createdAt:    createdAt.UTC(),
		modifiedAt:   modifiedAt.UTC(),
		name:         name,
		description:  description,
		contactCount: contactCount,
	}
}

func NewUpdater(id uuid.UUID, name name.Name, description description.Description) Updater {
	return &group{
		id:          id,
		modifiedAt:  time.Now().UTC(),
		name:        name,
		description: description,
	}
}

func NewDeleter(id uuid.UUID) Deleter {
	return &group{
		id:         id,
		modifiedAt: time.Now().UTC(),
	}
}
