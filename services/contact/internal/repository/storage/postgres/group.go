package postgres

import (
	"database/sql"
	"errors"
	"time"

	"github.com/Masterminds/squirrel"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"

	"app/pkg/tools/converter"
	"app/pkg/tools/transaction"
	"app/pkg/type/columnCode"
	"app/pkg/type/context"
	log "app/pkg/type/logger"
	"app/pkg/type/queryParameter"
	"app/services/contact/internal/domain/group"
	"app/services/contact/internal/repository/storage_template_1/postgres/dao"
	"app/services/contact/internal/useCase"
)

var mappingSortGroup = map[columnCode.ColumnCode]string{
	"id":           "id",
	"name":         "name",
	"description":  "description",
	"contactCount": "contact_count",
}

func (r *Repository) CreateGroup(c context.Context, creator group.Creator) (group.Reader, error) {

	ctx := c.CopyWithTimeout(r.options.Timeout)
	defer ctx.Cancel()

	query, args, err := r.genSQL.Insert("slurm.group").
		Columns(
			"id",
			"name",
			"description",
			"created_at",
			"modified_at",
		).
		Values(
			creator.ID(),
			creator.Name().Value(),
			creator.Description().Value(),
			creator.CreatedAt(),
			creator.ModifiedAt()).
		Suffix(`RETURNING
			id,
			name,
			description,
			created_at,
			modified_at,
			contact_count
			`,
		).
		ToSql()
	if err != nil {
		return nil, log.ErrorWithContext(ctx, err)
	}

	rows, err := r.db.Query(ctx, query, args...)
	if err != nil {
		return nil, log.ErrorWithContext(ctx, err)
	}

	var daoGroup []*dao.Group
	if err = pgxscan.ScanAll(&daoGroup, rows); err != nil {
		return nil, log.ErrorWithContext(ctx, err)
	}

	if len(daoGroup) == 0 {
		return nil, errors.New("group not found")
	}

	return daoGroup[0].ToDomainGroup()
}

func (r *Repository) UpdateGroup(c context.Context, updater group.Updater) (group.Reader, error) {

	ctx := c.CopyWithTimeout(r.options.Timeout)
	defer ctx.Cancel()

	query, args, err := r.genSQL.Update("slurm.group").
		Set("name", updater.Name().Value()).
		Set("description", updater.Description().Value()).
		Set("modified_at", updater.ModifiedAt()).
		Where(squirrel.And{
			squirrel.Eq{
				"id":          updater.ID().String(),
				"is_archived": false,
			},
		}).
		Suffix(`RETURNING
			id,
			name,
			description,
			created_at,
			modified_at,
			contact_count`,
		).
		ToSql()
	if err != nil {
		return nil, log.ErrorWithContext(ctx, err)
	}

	rows, err := r.db.Query(ctx, query, args...)
	if err != nil {
		return nil, log.ErrorWithContext(ctx, err)
	}

	var daoGroup []*dao.Group
	if err = pgxscan.ScanAll(&daoGroup, rows); err != nil {
		return nil, log.ErrorWithContext(ctx, err)
	}

	if len(daoGroup) == 0 {
		return nil, errors.New("group not found")
	}

	return daoGroup[0].ToDomainGroup()
}

func (r *Repository) DeleteGroup(c context.Context, deleter group.Deleter) error {

	ctx := c.CopyWithTimeout(r.options.Timeout)
	defer ctx.Cancel()

	tx, err := r.db.Begin(ctx)
	if err != nil {
		return log.ErrorWithContext(ctx, err)
	}

	defer func(ctx context.Context, t pgx.Tx) {
		err = transaction.Finish(ctx, t, err)
	}(ctx, tx)

	if err = r.deleteGroupTx(ctx, tx, deleter.ID()); err != nil {
		return err
	}

	return nil
}

func (r *Repository) deleteGroupTx(ctx context.Context, tx pgx.Tx, ID uuid.UUID) error {
	query, args, err := r.genSQL.Update("slurm.group").
		Set("is_archived", true).
		Set("modified_at", time.Now().UTC()).
		Where(squirrel.Eq{
			"id":          ID,
			"is_archived": false,
		}).ToSql()

	if err != nil {
		return log.ErrorWithContext(ctx, err)
	}

	if _, err = tx.Exec(ctx, query, args...); err != nil {
		return log.ErrorWithContext(ctx, err)
	}

	if err = r.clearGroupTx(ctx, tx, ID); err != nil {
		return err
	}

	return nil
}

func (r *Repository) clearGroupTx(ctx context.Context, tx pgx.Tx, groupID uuid.UUID) error {
	query, args, err := r.genSQL.
		Delete("slurm.contact_in_group").
		Where(squirrel.Eq{"group_id": groupID}).
		ToSql()
	if err != nil {
		return log.ErrorWithContext(ctx, err)
	}

	if _, err = tx.Exec(ctx, query, args...); err != nil {
		return log.ErrorWithContext(ctx, err)
	}

	if err = r.updateGroupContactCount(ctx, tx, groupID); err != nil {
		return err
	}

	return nil
}

func (r *Repository) ListGroup(c context.Context, parameter queryParameter.QueryParameter) ([]group.Reader, error) {

	ctx := c.CopyWithTimeout(r.options.Timeout)
	defer ctx.Cancel()

	tx, err := r.db.Begin(ctx)
	if err != nil {
		return nil, log.ErrorWithContext(ctx, err)
	}

	defer func(ctx context.Context, t pgx.Tx) {
		err = transaction.Finish(ctx, t, err)
	}(ctx, tx)

	response, err := r.listGroupTx(ctx, tx, parameter)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (r *Repository) listGroupTx(ctx context.Context, tx pgx.Tx, parameter queryParameter.QueryParameter) ([]group.Reader, error) {
	var result []group.Reader

	var builder = r.genSQL.Select(
		"id",
		"name",
		"description",
		"created_at",
		"modified_at",
		"contact_count",
		"is_archived",
	).
		From("slurm.group")

	builder = builder.Where(squirrel.Eq{"is_archived": false})

	if len(parameter.Sorts) > 0 {
		builder = builder.OrderBy(parameter.Sorts.Parsing(mappingSortGroup)...)
	} else {
		builder = builder.OrderBy("created_at DESC")
	}

	if parameter.Pagination.Limit > 0 {
		builder = builder.Limit(parameter.Pagination.Limit)
	}
	if parameter.Pagination.Offset > 0 {
		builder = builder.Offset(parameter.Pagination.Offset)
	}

	query, args, err := builder.ToSql()
	rows, err := tx.Query(ctx, query, args...)
	if err != nil {
		return nil, log.ErrorWithContext(ctx, err)
	}

	var groups []*dao.Group
	if err = pgxscan.ScanAll(&groups, rows); err != nil {
		return nil, log.ErrorWithContext(ctx, err)
	}

	for _, g := range groups {
		domainGroup, err := g.ToDomainGroup()
		if err != nil {
			return nil, log.ErrorWithContext(ctx, err)
		}
		result = append(result, domainGroup)
	}
	return result, nil
}

func (r *Repository) ReadGroupByID(c context.Context, ID uuid.UUID) (group.Reader, error) {

	ctx := c.CopyWithTimeout(r.options.Timeout)
	defer ctx.Cancel()

	tx, err := r.db.Begin(ctx)
	if err != nil {
		return nil, log.ErrorWithContext(ctx, err)
	}

	defer func(ctx context.Context, t pgx.Tx) {
		err = transaction.Finish(ctx, t, err)
	}(ctx, tx)

	response, err := r.oneGroupTx(ctx, tx, ID)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (r *Repository) oneGroupTx(ctx context.Context, tx pgx.Tx, ID uuid.UUID) (response group.Reader, err error) {

	var builder = r.genSQL.Select(
		"id",
		"name",
		"description",
		"created_at",
		"modified_at",
		"contact_count",
		"is_archived",
	).
		From("slurm.group")

	builder = builder.Where(squirrel.Eq{"is_archived": false, "id": ID})

	query, args, err := builder.ToSql()
	rows, err := tx.Query(ctx, query, args...)
	if err != nil {
		return nil, log.ErrorWithContext(ctx, err)
	}

	var daoGroup []*dao.Group
	if err = pgxscan.ScanAll(&daoGroup, rows); err != nil {
		return nil, log.ErrorWithContext(ctx, err)
	}

	if len(daoGroup) == 0 {
		return nil, useCase.ErrGroupNotFound
	}

	return daoGroup[0].ToDomainGroup()
}

func (r *Repository) CountGroup(ctx context.Context) (uint64, error) {
	var builder = r.genSQL.Select(
		"COUNT(id)",
	).From("slurm.group")

	builder = builder.Where(squirrel.Eq{"is_archived": false})

	query, args, err := builder.ToSql()
	if err != nil {
		return 0, log.ErrorWithContext(ctx, err)
	}

	row := r.db.QueryRow(ctx, query, args...)
	var total uint64

	if err = row.Scan(&total); err != nil {
		return 0, log.ErrorWithContext(ctx, err)
	}

	return total, nil
}

func (r *Repository) updateGroupsContactCountByFilters(ctx context.Context, tx pgx.Tx, ID uuid.UUID) error {

	builder := r.genSQL.Select("contact_in_group.group_id").
		From("slurm.contact_in_group").
		InnerJoin("slurm.contact ON contact_in_group.contact_id = contact.id").
		GroupBy("contact_in_group.group_id")

	builder = builder.Where(squirrel.Eq{"contact_in_group.contact_id": ID})

	query, args, err := builder.ToSql()
	if err != nil {
		return log.ErrorWithContext(ctx, err)
	}

	rows, err := tx.Query(ctx, query, args...)
	if err != nil {
		return log.ErrorWithContext(ctx, err)
	}
	var groupIDs []uuid.UUID
	for rows.Next() {
		var groupID sql.NullString
		if err = rows.Scan(&groupID); err != nil {
			return log.ErrorWithContext(ctx, err)
		}
		groupIDs = append(groupIDs, converter.StringToUUID(groupID.String))
	}

	for _, groupID := range groupIDs {
		if err = r.updateGroupContactCount(ctx, tx, groupID); err != nil {
			return err
		}
	}

	if err = rows.Err(); err != nil {
		return log.ErrorWithContext(ctx, err)
	}

	return nil
}

func (r *Repository) updateGroupContactCount(ctx context.Context, tx pgx.Tx, groupID uuid.UUID) error {
	subSelect := r.genSQL.Select("count(contact_in_group.id)").
		From("slurm.contact_in_group").
		InnerJoin("slurm.contact ON contact_in_group.contact_id = contact.id").
		Where(squirrel.Eq{"group_id": groupID, "is_archived": false})

	query, _, err := r.genSQL.
		Update("slurm.group").
		Set("contact_count", subSelect).
		Where(squirrel.Eq{"id": groupID}).
		ToSql()
	if err != nil {
		return log.ErrorWithContext(ctx, err)
	}

	var args = []interface{}{groupID, false}

	if _, err = tx.Exec(ctx, query, args...); err != nil {
		return log.ErrorWithContext(ctx, err)
	}
	return nil
}
