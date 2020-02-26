package models

import (
	"github.com/lucasmbaia/punctorum/api/repository/filter"
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

func (u *Users) Get(filters []filter.Filters, args ...interface{}) (users []UsersFields, err error) {
	err = u.DB.Read(filters, &users, args)
	return
}

func (u *Users) Post(data *UsersFields) (async bool, err error) {
	data.ID = "9876"
	return
}
