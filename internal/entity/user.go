package entity

type User struct {
	Entity `json:","`

	Name     string  `json:"name,omitempty"`
	Handle   string  `json:"handle,omitempty"`
	Email    *string `json:"-"`
	Password string  `json:"-"`
}
