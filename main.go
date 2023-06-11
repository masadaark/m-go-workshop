package main

import (
	"gin-intro/api"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	authenApi := router.Group("/authen")
	{
		authenApi.GET("/login", api.Login)
		authenApi.GET("/register", api.Register)
	}
	router.Run(":899")
}
