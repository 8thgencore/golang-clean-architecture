package group

import (
	"github.com/google/uuid"

	"app/pkg/type/context"
	"app/pkg/type/queryParameter"
	"app/services/contact/internal/domain/group"
)

func (uc *UseCase) Create(ctx context.Context, creator group.Creator) (group.Reader, error) {
	return uc.adapterStorage.CreateGroup(ctx, creator)
}

func (uc *UseCase) Update(ctx context.Context, updater group.Updater) (group.Reader, error) {
	return uc.adapterStorage.UpdateGroup(ctx, updater)
}

func (uc *UseCase) Delete(ctx context.Context, deleter group.Deleter) error {
	return uc.adapterStorage.DeleteGroup(ctx, deleter)
}

func (uc *UseCase) List(ctx context.Context, parameter queryParameter.QueryParameter) ([]group.Reader, error) {
	return uc.adapterStorage.ListGroup(ctx, parameter)
}

func (uc *UseCase) ReadByID(ctx context.Context, ID uuid.UUID) (group.Reader, error) {
	return uc.adapterStorage.ReadGroupByID(ctx, ID)
}

func (uc *UseCase) Count(ctx context.Context) (uint64, error) {
	return uc.adapterStorage.CountGroup(ctx)
}
