package domain

import (
	"fmt"
	"net/http"
)

type ErrorTag int

const (
	ETNotFound ErrorTag = iota
	ETAccessDenied
	ETAuthRequired
	ETConflict
	ETBadRequest
)

type TaggedError struct {
	Tag ErrorTag
	Msg string
}

func (e TaggedError) Error() string {
	return e.Msg
}

func (t ErrorTag) ToString() string {
	switch t {
	case ETNotFound:
		return "Not Found"
	case ETAccessDenied:
		return "Access Denied"
	case ETAuthRequired:
		return "Auth Required"
	case ETConflict:
		return "Conflict"
	case ETBadRequest:
		return "Bad Request"
	}
	return "Invalid Error Tag"
}

func (t ErrorTag) ToHTTPStatus() int {
	switch t {
	case ETNotFound:
		return http.StatusNotFound
	case ETAccessDenied:
		return http.StatusForbidden
	case ETAuthRequired:
		return http.StatusUnauthorized
	case ETConflict:
		return http.StatusConflict
	case ETBadRequest:
		return http.StatusBadRequest
	}
	return http.StatusInternalServerError
}

func ErrorNotFound(format string, args ...any) error {
	if len(args) > 0 {
		format = fmt.Sprintf(format, args)
	}
	return TaggedError{
		Tag: ETNotFound,
		Msg: format,
	}
}

func ErrorAccessDenied(format string, args ...any) error {
	if len(args) > 0 {
		format = fmt.Sprintf(format, args)
	}
	return TaggedError{
		Tag: ETAccessDenied,
		Msg: format,
	}
}

func ErrorAuthRequired(format string, args ...any) error {
	if len(args) > 0 {
		format = fmt.Sprintf(format, args)
	}
	return TaggedError{
		Tag: ETAuthRequired,
		Msg: format,
	}
}

func ErrorConflict(format string, args ...any) error {
	if len(args) > 0 {
		format = fmt.Sprintf(format, args)
	}
	return TaggedError{
		Tag: ETConflict,
		Msg: format,
	}
}

func ErrorBadRequest(format string, args ...any) error {
	if len(args) > 0 {
		format = fmt.Sprintf(format, args)
	}
	return TaggedError{
		Tag: ETBadRequest,
		Msg: format,
	}
}
