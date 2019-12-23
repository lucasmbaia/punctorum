package main

import (
	//"github.com/lucasmbaia/punctorum/api/models/interfaces"
	"github.com/gin-gonic/gin"
	//"github.com/lucasmbaia/punctorum/api/models"
	"github.com/lucasmbaia/punctorum/api/controllers"
)

func main() {
	var (
		g     *gin.Engine
		users *controllers.Users
	)

	users = controllers.NewUsers()
	g = gin.Default()

	v1 := g.Group("/teste/:ID/users")
	{
		v1.GET("", users.Get)
		v1.POST("", users.Post)
	}

	g.Run()
}
