package entity

type User struct {
	Entity `json:","`

	Name     string `json:"name"`
	Login    string `json:"-"`
	Password string `json:"-"`
}
