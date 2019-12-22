package main

import (
	//"github.com/lucasmbaia/punctorum/api/models/interfaces"
	"github.com/lucasmbaia/punctorum/api/models"
	"github.com/lucasmbaia/punctorum/api/controllers"
)

func main() {
	var users = controllers.NewUsers()
	var data interface{}

	data = models.UsersFields{
		Name:	"lucas",
	}

	users.Get(data)
	
	//var m interfaces.Models = models.NewUsers()

	//m.Get(data)
}
