package entity

import (
	"github.com/google/uuid"
)

type User struct {
	*ObjectStatus `json:"status,omitempty"`

	ID uuid.UUID `json:"id"`

	*UserAuthData `json:"-"`
	*UserData     `json:",omitempty"`
}

type UserAuthData struct {
	Login    string `json:"-"`
	Password string `json:"-"`
}

