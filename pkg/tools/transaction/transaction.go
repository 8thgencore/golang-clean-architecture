package transaction

import (
	"app/pkg/type/context"
	"app/pkg/type/logger"

	"github.com/jackc/pgx/v4"
)

func Finish(ctx context.Context, tx pgx.Tx, err error) error {
	if err != nil {
		if rollbackErr := tx.Rollback(ctx); rollbackErr != nil {
			return logger.ErrorWithContext(ctx, rollbackErr)
		}
		return err
	}
	if commitErr := tx.Commit(ctx); commitErr != nil {

		return logger.ErrorWithContext(ctx, commitErr)
	}
	return nil

}
