package dao

import (
	"time"

	"github.com/google/uuid"

	"app/services/contact/internal/domain/group"
	"app/services/contact/internal/domain/group/description"
	"app/services/contact/internal/domain/group/name"
)

type Group struct {
	ID           uuid.UUID `db:"id"`
	Name         string    `db:"name"`
	Description  string    `db:"description"`
	CreatedAt    time.Time `db:"created_at"`
	ModifiedAt   time.Time `db:"modified_at"`
	ContactCount uint64    `db:"contact_count"`
	IsArchived   bool      `db:"is_archived"`
}

func (g *Group) ToDomainGroup() (group.Reader, error) {
	gN, err := name.New(g.Name)
	if err != nil {
		return nil, err
	}

	gD, err := description.New(g.Description)
	if err != nil {
		return nil, err
	}
	return group.NewReader(
		g.ID,
		g.CreatedAt,
		g.ModifiedAt,
		gN,
		gD,
		g.ContactCount,
	), nil
}
