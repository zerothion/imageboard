package rest

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/zerothion/imageboard/internal/delivery"
	"github.com/zerothion/imageboard/internal/domain"
	"github.com/zerothion/imageboard/internal/entity"
)

type userHandlers struct {
	userService domain.UserService
}

func AddUserHandlers(s *delivery.ServeMuxWrapper, userService domain.UserService) {
	h := userHandlers{userService}

	s.Handle("GET /api/users", delivery.NotImplementedHandler)
	s.Handle("GET /api/user/{id}", h.GetUserById)
	s.Handle("POST /api/user", h.CreateUser)
	s.Handle("DELETE /api/user/{id}", h.DeleteUser)
	s.Handle("PATCH /api/user/{id}", delivery.NotImplementedHandler)
}

func (h *userHandlers) GetUserById(w http.ResponseWriter, r *http.Request) error {
	id, err := uuid.Parse(r.PathValue("id"))
	if err != nil {
		return domain.ErrorBadRequest("Failed to parse user id - %s", err.Error())
	}
	user, err := h.userService.GetByID(r.Context(), id)
	if err != nil {
		return err
	}
	user_json, _ := json.Marshal(user)
	w.Write(user_json)
	return nil
}

func (h *userHandlers) CreateUser(w http.ResponseWriter, r *http.Request) error {
	login := r.PostFormValue("login")
	if login == "" {
		return domain.ErrorBadRequest("`login` is required for creating a new user")
	}

	password := r.PostFormValue("password")
	if password == "" {
		return domain.ErrorBadRequest("`password` is required for creating a new user")
	}

	name := r.PostFormValue("name")
	if name == "" {
		name = login
	}

	user := entity.User{
		Name:     name,
		Login:    login,
		Password: password,
	}
	err := h.userService.Create(r.Context(), &user)
	if err != nil {
		return err
	}

	result, err := json.Marshal(user)
	if err != nil {
		return err
	}

	w.WriteHeader(http.StatusCreated)
	w.Write(result)
	return nil
}

func (h *userHandlers) DeleteUser(w http.ResponseWriter, r *http.Request) error {
	id := r.PathValue("id")
	uid, err := uuid.Parse(id)
	if err != nil {
		return domain.ErrorBadRequest("Failed to parse user id - %s", err.Error())
	}
	return h.userService.Delete(r.Context(), uid)
}
