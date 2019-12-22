package controllers

import (
	"github.com/lucasmbaia/punctorum/api/models/interfaces"
	"github.com/lucasmbaia/punctorum/api/models"
)

type Users struct {
	Resources
}

func NewUsers() *Users {
	return &Users{
		Resources{
			GetModel: func() interfaces.Models {
				return models.NewResources(models.NewUsers())
			},
			GetFields: func() interface{} {
				return &models.Users{}
			},
		},
	}
}
