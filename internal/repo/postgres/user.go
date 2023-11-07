package postgres

import (
	"context"
	"errors"

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

func (u *userRepo) GetById(ctx context.Context, id uuid.UUID) (entity.User, error) {
	row := u.db.QueryRow(
		ctx, `
        SELECT
            e.created_at, e.deleted_at,
            u.user_id, u.name,
            u.login, u.password
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
		&user.Login, &user.Password,
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
            RETURNING entity_id
    `)
	err = row.Scan(&user.ID)
	if err != nil {
		return err
	}

	_, err = tx.Exec(
		ctx, `
        INSERT INTO users (user_id, name, login, password)
            VALUES ($1, $2, $3, $4)`,
		user.ID,
		user.Name,
		user.Login,
		user.Password,
	)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" { // 23505 - Unique Violation
			return domain.ErrorConflict("Login is already taken!")
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
