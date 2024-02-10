package domain

import (
	"context"
	"time"

	"github.com/google/uuid"

	"github.com/zerothion/imageboard/internal/entity"
	"github.com/zerothion/imageboard/internal/repo"
)

type UserService interface {
	Fetch(ctx context.Context, before time.Time, limit uint64, offset uint64) ([]entity.User, error)
	GetByID(ctx context.Context, id uuid.UUID) (entity.User, error)
	Create(ctx context.Context, user *entity.User) error
	Delete(ctx context.Context, id uuid.UUID) error
}

type userService struct {
	userRepo repo.UserRepo
}

func NewUserService(userRepo repo.UserRepo) UserService {
	return &userService{
		userRepo,
	}
}

func (s *userService) Fetch(ctx context.Context, before time.Time, limit uint64, offset uint64) ([]entity.User, error) {
	if limit > 200 {
		limit = 200
	}
	return s.userRepo.Fetch(ctx, before, limit, offset)
}

func (s *userService) GetByID(ctx context.Context, id uuid.UUID) (entity.User, error) {
	return s.userRepo.GetById(ctx, id)
}

func (s *userService) Create(ctx context.Context, user *entity.User) error {
	if len(user.Login) < 3 {
		return ErrorUnprocessableContent("Login must be atleast 3 character long")
	}

	salt, err := generateSalt()
	if err != nil {
		return err
	}
	user.Password = hashPassword([]byte(user.Password), salt)

	err = s.userRepo.Store(ctx, user)
	return err
}

func (s *userService) Delete(ctx context.Context, id uuid.UUID) error {
    return s.userRepo.Delete(ctx, id)
}
