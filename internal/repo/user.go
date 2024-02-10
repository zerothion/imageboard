package repo

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/zerothion/imageboard/internal/entity"
)

type UserRepo interface {
	Fetch(ctx context.Context, before time.Time, limit uint64, offset uint64) ([]entity.User, error)
	GetById(ctx context.Context, id uuid.UUID) (entity.User, error)
	Store(ctx context.Context, user *entity.User) error
	Update(ctx context.Context, user *entity.User) error
	Delete(ctx context.Context, id uuid.UUID) error
}
