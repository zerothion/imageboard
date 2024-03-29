package postgres

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/zerothion/imageboard/internal/domain"
	"github.com/zerothion/imageboard/internal/entity"
	"github.com/zerothion/imageboard/internal/repo"
)

func NewUserRepo(db *pgxpool.Pool) repo.UserRepo {
	return &userRepo{
		db: db,
	}
}

type userRepo struct {
	db *pgxpool.Pool
}

func (u *userRepo) Fetch(ctx context.Context, before time.Time, limit uint64, offset uint64) ([]entity.User, error) {
	rows, err := u.db.Query(
		ctx, `
        SELECT
            e.created_at, e.deleted_at,
            u.user_id, u.name,
            u.handle
        FROM users u
        INNER JOIN entities e
            ON u.user_id = e.entity_id
        WHERE e.created_at <= $1 AND (e.deleted_at IS NULL OR e.deleted_at > $1)
        ORDER BY e.created_at DESC
        LIMIT $2
        OFFSET $3
        `, before, limit, offset,
	)
	if err != nil {
		return nil, err
	}

	users := make([]entity.User, 0)
	for rows.Next() {
		var user entity.User
		err = rows.Scan(
			&user.CreatedAt, &user.DeletedAt,
			&user.ID, &user.Name,
			&user.Handle,
		)
		if err != nil {
			return users, err
		}
		users = append(users, user)
	}

	return users, nil
}

func (u *userRepo) GetById(ctx context.Context, id uuid.UUID) (entity.User, error) {
	row := u.db.QueryRow(
		ctx, `
        SELECT
            e.created_at, e.deleted_at,
            u.user_id, u.name,
            u.email,
            u.handle, u.password
        FROM users u
        INNER JOIN entities e
            ON u.user_id = e.entity_id
        WHERE u.user_id = $1`,
		id,
	)

	var user entity.User
	err := row.Scan(
		&user.CreatedAt, &user.DeletedAt,
		&user.ID, &user.Name,
		&user.Email,
		&user.Handle, &user.Password,
	)
	return user, err
}

func (u *userRepo) Store(ctx context.Context, user *entity.User) error {
	tx, err := u.db.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return err
	}

	row := tx.QueryRow(ctx, `
        INSERT INTO entities
            DEFAULT VALUES
            RETURNING entity_id, created_at
    `)
	err = row.Scan(&user.ID, &user.CreatedAt)
	if err != nil {
		return err
	}

	_, err = tx.Exec(
		ctx, `
        INSERT INTO users (user_id, name, email, handle, password)
            VALUES ($1, $2, $3, $4, $5)`,
		user.ID,
		user.Name,
		user.Email,
		user.Handle,
		user.Password,
	)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" { // 23505 - Unique Violation
			return domain.ErrorConflict("The handle @%s is already taken!", user.Handle)
		}
		return err
	}

	return tx.Commit(ctx)
}

func (u *userRepo) Update(ctx context.Context, user *entity.User) error {
	panic("not implemented") // TODO: Implement
}

func (u *userRepo) Delete(ctx context.Context, id uuid.UUID) error {
	c, err := u.db.Exec(ctx, `
        UPDATE entities
            SET deleted_at = now()
        WHERE entity_id = $1 AND deleted_at IS NULL
    `, id)
	if err != nil {
		return err
	}
	if c.RowsAffected() == 0 {
		return domain.ErrorNotFound("User not found")
	}
	return nil
}
