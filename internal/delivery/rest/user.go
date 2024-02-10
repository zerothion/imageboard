package rest

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

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

	s.Handle("GET /api/users", h.ListUsers)
	s.Handle("GET /api/user/{id}", h.GetUserById)
	s.Handle("POST /api/user", h.CreateUser)
	s.Handle("DELETE /api/user/{id}", h.DeleteUser)
	s.Handle("PATCH /api/user/{id}", delivery.NotImplementedHandler)
}

func (h *userHandlers) ListUsers(w http.ResponseWriter, r *http.Request) error {
	var err error
	offset := uint64(0)
	if raw := r.URL.Query().Get("offset"); raw != "" {
		offset, err = strconv.ParseUint(raw, 10, 64)
		if err != nil {
			return domain.ErrorBadRequest("Failed to parse `offset` - must be a valid uint64")
		}
	}

	limit := uint64(50)
	if raw := r.URL.Query().Get("limit"); raw != "" {
		limit, err = strconv.ParseUint(raw, 10, 64)
		if err != nil {
			return domain.ErrorBadRequest("Failed to parse `limit` - must be a valid uint64")
		}
	}

	before := time.Now()
	if raw := r.URL.Query().Get("before"); raw != "" {
		val, err := strconv.ParseInt(raw, 10, 64)
		if err != nil {
			return domain.ErrorBadRequest("Failed to parse `before` - must be a unix timestamp (int64)")
		}
		before = time.Unix(val, 0)
	}

	users, err := h.userService.Fetch(r.Context(), before, limit, offset)
	if err != nil {
		return err
	}

	result, err := json.Marshal(users)
	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(result)
	return nil
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

	result, err := json.Marshal(user)
	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(result)
	return nil
}

func (h *userHandlers) CreateUser(w http.ResponseWriter, r *http.Request) error {
	handle := r.PostFormValue("handle")
	if handle == "" {
		return domain.ErrorBadRequest("`handle` is required for creating a new user")
	}

	password := r.PostFormValue("password")
	if password == "" {
		return domain.ErrorBadRequest("`password` is required for creating a new user")
	}

	name := r.PostFormValue("name")

	var email *string
	if val := r.PostFormValue("email"); val != "" {
		email = &val
		// todo: validate email?
	}

	user := entity.User{
		Name:     name,
		Email:    email,
		Handle:   handle,
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

	w.Header().Set("Content-Type", "application/json")
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
