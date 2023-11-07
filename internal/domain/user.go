package domain

import (
	"context"

	"github.com/google/uuid"

	"github.com/zerothion/imageboard/internal/entity"
	"github.com/zerothion/imageboard/internal/repo"
)

type UserService interface {
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

func (s *userService) GetByID(ctx context.Context, id uuid.UUID) (entity.User, error) {
	return s.userRepo.GetById(ctx, id)
}

func (s *userService) Create(ctx context.Context, user *entity.User) error {
	if len(user.Login) < 3 {
		return ErrorBadRequest("Login must be atleast 3 character long")
	}

	err = s.userRepo.Store(ctx, user)
	return err
}

func (s *userService) Delete(ctx context.Context, id uuid.UUID) error {
    return s.userRepo.Delete(ctx, id)
}
