package main

import (
	"gin-intro/api"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	api.Setup(router)

	router.Run(":899")
}
