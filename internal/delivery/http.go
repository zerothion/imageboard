package delivery

import (
	"github.com/go-chi/chi/v5"
)

type ServerHTTP struct {
	chi.Router
}

func NewHTTP() *ServerHTTP {
	return &ServerHTTP{chi.NewRouter()}
}

func (s *ServerHTTP) Post(pattern string, handler Handler) {
	s.Router.Post(pattern, handler.AsHandlerFunc())
}

func (s *ServerHTTP) Delete(pattern string, handler Handler) {
	s.Router.Delete(pattern, handler.AsHandlerFunc())
}

func (s *ServerHTTP) Put(pattern string, handler Handler) {
	s.Router.Put(pattern, handler.AsHandlerFunc())
}

func (s *ServerHTTP) Get(pattern string, handler Handler) {
	s.Router.Get(pattern, handler.AsHandlerFunc())
}

func (s *ServerHTTP) Patch(pattern string, handler Handler) {
	s.Router.Patch(pattern, handler.AsHandlerFunc())
}
