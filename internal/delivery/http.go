package delivery

import (
	"net/http"
)

type ServeMuxWrapper struct {
	*http.ServeMux
}

func NewServeMux() *ServeMuxWrapper {
	return &ServeMuxWrapper{
		http.NewServeMux(),
	}
}

func (s ServeMuxWrapper) Handle(pattern string, handler Handler) {
	s.ServeMux.Handle(pattern, handler)
}
