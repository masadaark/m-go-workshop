package api

import "github.com/gin-gonic/gin"

func Setup(rounter *gin.Engine) {

	authenApi := rounter.Group("/authen")
	{
		authenApi.GET("/login", Login)
		authenApi.GET("/register", Register)
	}

	returnApi := rounter.Group("/json")
	{
		returnApi.GET("/someJSON", SomeJSON)
		returnApi.GET("/moreJSON", MoreJSON)
		returnApi.GET("/someXML", SomeXML)
		returnApi.GET("/someYAML", SomeYAML)
	}

	imageApi := rounter.Group("/image")
	{
		imageApi.GET("/jung", Jung)

	}

	publicApi := rounter.Group("/public")
	{
		publicApi.Static("/jung", "public")
	}
}
