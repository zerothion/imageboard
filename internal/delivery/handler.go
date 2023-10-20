package delivery

import (
	"log/slog"
	"net/http"

	"github.com/zerothion/imageboard/internal/domain"
)

type Handler func(w http.ResponseWriter, r *http.Request) error

func (h Handler) AsHandlerFunc() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := h(w, r); err != nil {
			if te, ok := err.(domain.TaggedError); ok {
				w.WriteHeader(te.Tag.ToHTTPStatus())
				w.Write([]byte(te.Msg)) // todo: <- improve error message
			} else {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte("Internal Server Error"))
				// todo: change log args to contain more useful information about request
				slog.Error(
					"Handler failed a request",
					"err", err,
					"method", r.Method,
					"url", r.URL.Redacted(),
					"body", r.Body,
				)
			}
		}
	}
}
