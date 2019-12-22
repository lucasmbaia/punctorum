package models

import (
	"fmt"
)

type UsersFields struct {
	Name	string	`json:",omitempty"`
}

type Users struct {
	Resources
}

func NewUsers() *Users {
	var users = &Users{}

	return users
}

func (u *Users) Get(data interface{}) (users []UsersFields, err error) {
	fmt.Println("users")
	return
}
