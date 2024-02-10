package repo

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/zerothion/imageboard/internal/entity"
)

var (
	ErrStubUsed = errors.New("stub; no implementation")
)

// --- repo.UserRepo ---
type userRepoStub struct{}

func (*userRepoStub) Fetch(ctx context.Context, before time.Time, limit uint64, offset uint64) ([]entity.User, error) {
	return nil, fmt.Errorf("UserRepo.Fetch %w", ErrStubUsed)
}

func (r *userRepoStub) GetById(ctx context.Context, id uuid.UUID) (entity.User, error) {
	return entity.User{}, fmt.Errorf("UserRepo.GetById %w", ErrStubUsed)
}

func (r *userRepoStub) Store(ctx context.Context, user *entity.User) error {
	return fmt.Errorf("UserRepo.Store %w", ErrStubUsed)
}

func (r *userRepoStub) Update(ctx context.Context, user *entity.User) error {
	return fmt.Errorf("UserRepo.Update %w", ErrStubUsed)
}

func (r *userRepoStub) Delete(ctx context.Context, id uuid.UUID) error {
	return fmt.Errorf("UserRepo.Delete %w", ErrStubUsed)
}

func NewUserRepoStub() UserRepo {
	return &userRepoStub{}
}
