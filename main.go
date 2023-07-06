package main

import (
	"gin-intro/api"
	"gin-intro/pgSql"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	pgSql.Main()

	api.Setup(router)

	router.Run(":899")

}
