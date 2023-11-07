package delivery

import (
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/zerothion/imageboard/internal/delivery/utils"
	"github.com/zerothion/imageboard/internal/domain"
	"github.com/zerothion/imageboard/internal/repo"
)

const DefaultAddress = ":80"

type ServerHTTP struct {
	Repos

	router      chi.Router
	UserService domain.UserService
}

type Repos struct {
	UserRepo  repo.UserRepo
}

func NewHTTP(repos Repos) *ServerHTTP {
	s := &ServerHTTP{}
	s.router = chi.NewRouter()

	if utils.IsAnyFieldNil(repos) {
		slog.Error("All repos must be set to non-nil values!")
	}

	s.Repos = repos
	s.UserService = domain.NewUserService(s.UserRepo)

	return s
}

func (s *ServerHTTP) Post(pattern string, handler Handler) {
	s.router.Post(pattern, handler.AsHandlerFunc())
}

func (s *ServerHTTP) Delete(pattern string, handler Handler) {
	s.router.Delete(pattern, handler.AsHandlerFunc())
}

func (s *ServerHTTP) Put(pattern string, handler Handler) {
	s.router.Put(pattern, handler.AsHandlerFunc())
}

func (s *ServerHTTP) Get(pattern string, handler Handler) {
	s.router.Get(pattern, handler.AsHandlerFunc())
}

func (s *ServerHTTP) Patch(pattern string, handler Handler) {
	s.router.Patch(pattern, handler.AsHandlerFunc())
}

func (s *ServerHTTP) Serve(addr string) error {
	slog.Info("Listening for HTTP", "addr", addr)
	return http.ListenAndServe(addr, s.router)
}

func (s *ServerHTTP) ServeDefault() error {
	return s.Serve(DefaultAddress)
}
