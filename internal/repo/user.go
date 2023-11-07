package repo

import (
	"context"

	"github.com/google/uuid"
	"github.com/zerothion/imageboard/internal/entity"
)

type UserRepo interface {
	GetById(ctx context.Context, id uuid.UUID) (entity.User, error)
	Store(ctx context.Context, user *entity.User) error
	Update(ctx context.Context, user *entity.User) error
	Delete(ctx context.Context, id uuid.UUID) error
}
