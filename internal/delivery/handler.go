package delivery

import (
	"log/slog"
	"net/http"

	"github.com/zerothion/imageboard/internal/domain"
)

type Handler func(w http.ResponseWriter, r *http.Request) error

func NotImplementedHandler(w http.ResponseWriter, r *http.Request) error {
	w.WriteHeader(http.StatusNotImplemented)
	w.Write([]byte("This handler is yet to be implemented >_<"))
	slog.Warn("An unimplemented handler was called", "method", r.Method, "URL", r.URL.Redacted())
	return nil
}

func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
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
