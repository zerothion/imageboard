package domain

import (
	"fmt"
	"net/http"
)

type ErrorTag int

const (
	ETNotFound ErrorTag = iota
	ETForbidden
	ETUnauthorized
	ETConflict
	ETBadRequest
	ETUnprocessableContent
)

type TaggedError struct {
	Tag ErrorTag
	Msg string
}

func (e TaggedError) Error() string {
	return e.Tag.ToString() + ": " + e.Msg
}

func (t ErrorTag) ToString() string {
	switch t {
	case ETNotFound:
		return "Not Found"
	case ETForbidden:
		return "Forbidden"
	case ETUnauthorized:
		return "Unauthorized"
	case ETConflict:
		return "Conflict"
	case ETBadRequest:
		return "Bad Request"
	case ETUnprocessableContent:
		return "Unprocessable Content"
	}
	return "<Invalid Tag>"
}

func (t ErrorTag) ToHTTPStatus() int {
	switch t {
	case ETNotFound:
		return http.StatusNotFound
	case ETForbidden:
		return http.StatusForbidden
	case ETUnauthorized:
		return http.StatusUnauthorized
	case ETConflict:
		return http.StatusConflict
	case ETBadRequest:
		return http.StatusBadRequest
	case ETUnprocessableContent:
		return http.StatusUnprocessableEntity
	}
	return http.StatusInternalServerError
}

func ErrorNotFound(format string, args ...interface{}) error {
	if len(args) > 0 {
		format = fmt.Sprintf(format, args)
	}
	return TaggedError{
		Tag: ETNotFound,
		Msg: format,
	}
}

func ErrorForbidden(format string, args ...interface{}) error {
	if len(args) > 0 {
		format = fmt.Sprintf(format, args)
	}
	return TaggedError{
		Tag: ETForbidden,
		Msg: format,
	}
}

func ErrorUnauthorized(format string, args ...interface{}) error {
	if len(args) > 0 {
		format = fmt.Sprintf(format, args)
	}
	return TaggedError{
		Tag: ETUnauthorized,
		Msg: format,
	}
}

func ErrorConflict(format string, args ...interface{}) error {
	if len(args) > 0 {
		format = fmt.Sprintf(format, args)
	}
	return TaggedError{
		Tag: ETConflict,
		Msg: format,
	}
}

func ErrorBadRequest(format string, args ...interface{}) error {
	if len(args) > 0 {
		format = fmt.Sprintf(format, args)
	}
	return TaggedError{
		Tag: ETBadRequest,
		Msg: format,
	}
}

func ErrorUnprocessableContent(format string, args ...interface{}) error {
	if len(args) > 0 {
		format = fmt.Sprintf(format, args)
	}
	return TaggedError{
		Tag: ETUnprocessableContent,
		Msg: format,
	}
}
