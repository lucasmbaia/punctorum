package models

import (
)

type UsersFields struct {
	ID	string	`json:",omitempty"`
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
	users = []UsersFields{
		{ID:   "1234", Name: "lucas"},
	}

	return
}

func (u *Users) Post(data *UsersFields) (async bool, err error) {
	data.ID = "9876"
	return
}
