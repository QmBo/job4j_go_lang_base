package repository

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"job4j.ru/go-lang-base/internal/trackerstore"
)

type RepoPg struct {
	pool *pgxpool.Pool
}

func NewRepoPg(pool *pgxpool.Pool) *RepoPg {
	return &RepoPg{pool: pool}
}

func (r *RepoPg) Create(ctx context.Context, it trackerstore.Item) error {
	_, err := r.pool.Exec(
		ctx,
		`insert into items(id, name, position) values($1, $2, $3)`,
		it.ID, it.Name, it.Position,
	)
	if err != nil {
		return trackerstore.ErrRPoolExec(err)
	}
	return nil
}

func (r *RepoPg) List(ctx context.Context) ([]trackerstore.Item, error) {
	rows, err := r.pool.Query(ctx, `select id, name, position from items order by position`)
	if err != nil {
		return nil, trackerstore.ErrRPoolQuery(err)
	}
	return getList(rows)
}

func (r *RepoPg) Get(ctx context.Context, id string) (trackerstore.Item, error) {
	var it trackerstore.Item
	err := r.pool.QueryRow(
		ctx,
		`select id, name, position from items where id = $1`,
		id,
	).Scan(&it.ID, &it.Name, &it.Position)

	if err != nil {
		return it, trackerstore.ErrRPoolQueryRow(err)
	}
	return it, nil
}

func (r *RepoPg) GetLastPosition(ctx context.Context) (int, error) {
	var position int
	err := r.pool.QueryRow(
		ctx,
		`select position from items order by position desc limit 1`,
	).Scan(&position)

	if err != nil && errors.Is(err, pgx.ErrNoRows) {
		err = nil
	}

	if err != nil {
		return -1, trackerstore.ErrRPoolQueryRow(err)
	}

	return position, nil
}

func (r *RepoPg) Update(ctx context.Context, position int, item trackerstore.Item) error {
	_, err := r.pool.Exec(
		ctx,
		`update items set name = $2 where position = $1`,
		position, item.Name,
	)
	if err != nil {
		return trackerstore.ErrRPoolExec(err)
	}

	return nil
}

func (r *RepoPg) UpdateByID(ctx context.Context, item trackerstore.Item) error {
	_, err := r.pool.Exec(
		ctx,
		`update items set name = $2 where id = $1`,
		item.ID, item.Name,
	)
	if err != nil {
		return trackerstore.ErrRPoolExec(err)
	}

	return nil
}

func (r *RepoPg) Delete(ctx context.Context, position int) error {
	_, err := r.pool.Exec(
		ctx,
		`delete from items where position = $1`,
		position,
	)
	if err != nil {
		return trackerstore.ErrRPoolExec(err)
	}

	return nil
}

func (r *RepoPg) DeleteByID(ctx context.Context, id string) error {
	_, err := r.pool.Exec(
		ctx,
		`delete from items where id = $1`,
		id,
	)
	if err != nil {
		return trackerstore.ErrRPoolExec(err)
	}

	return nil
}

func (r *RepoPg) Find(ctx context.Context, name string) ([]trackerstore.Item, error) {
	rows, err := r.pool.Query(
		ctx,
		`select id, name, position from items where name like $1 order by position`,
		"%"+name+"%")
	if err != nil {
		return nil, trackerstore.ErrRPoolQuery(err)
	}
	return getList(rows)
}

func (r *RepoPg) Reorder(ctx context.Context, position int) error {
	rows, err := r.pool.Query(
		ctx,
		`select id, position from items where position > $1`,
		position,
	)
	if err != nil {
		return trackerstore.ErrRPoolQuery(err)
	}

	return r.changePositions(ctx, rows)
}

func (r *RepoPg) GetByPosition(ctx context.Context, position int) (trackerstore.Item, error) {
	var it trackerstore.Item
	err := r.pool.QueryRow(
		ctx,
		`select id, name, position from items where position = $1`,
		position,
	).Scan(&it.ID, &it.Name, &it.Position)
	if err != nil {
		return it, trackerstore.ErrRPoolQueryRow(err)
	}

	return it, nil
}

func (r *RepoPg) changePositions(ctx context.Context, rows pgx.Rows) error {
	defer rows.Close()

	for rows.Next() {
		var it trackerstore.Item
		if err := rows.Scan(&it.ID, &it.Position); err != nil {
			return trackerstore.ErrRRowScan(err)
		}
		_, err := r.pool.Exec(
			ctx,
			`update items set position = $1 where id = $2`,
			it.Position-1, it.ID,
		)
		if err != nil {
			return trackerstore.ErrRPoolExec(err)
		}
	}
	return nil
}

func getList(rows pgx.Rows) ([]trackerstore.Item, error) {
	defer rows.Close()

	var items []trackerstore.Item
	for rows.Next() {
		var item trackerstore.Item
		if err := rows.Scan(&item.ID, &item.Name, &item.Position); err != nil {
			return nil, trackerstore.ErrRRowScan(err)
		}
		items = append(items, item)
	}

	if err := rows.Err(); err != nil {
		return nil, trackerstore.ErrRRowError(err)
	}

	return items, nil
}
