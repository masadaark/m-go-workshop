package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default() //create router variable

	r.GET("/", func(c *gin.Context) {
		c.Data(http.StatusOK, "text/html; charset=uft-8", []byte("Root")) //(httpStatus,contentType,response)
	})

	r.GET("/test", func(c *gin.Context) {
		c.Data(200, "text/html; charset=uft-8", []byte("test"))
	})
	//query param http://localhost:85/login?username=hua&password=jai
	r.GET("/login", func(c *gin.Context) {
		username, password := c.Query("username"), c.Query("password")
		c.JSON(http.StatusOK, gin.H{"result": "ok", "username": username, "password": password})
	})

	r.GET("/stringnumbers", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"result": "ok", "username": "1000.00"})
	})

	r.POST("/login", postLogin)

	r.GET("/fullName/:firstName/:lastName", returnFullName)

	if err := r.Run(":85"); //run rounter (default port 8080)
	err != nil {
		log.Fatal(err)
	}
}

// http://localhost:85/fullName/Masada/Hiran
func returnFullName(c *gin.Context) {
	firstName, lastName := c.Param("firstName"), c.Param("lastName")
	c.JSON(http.StatusOK, gin.H{"fullName": firstName + " " + lastName})
}

func postLogin(c *gin.Context) {
	var form LoginFrom
	//if != nil error code
	if c.ShouldBind(&form) == nil {
		if form.Username == "admin" && form.Password == "1234" {
			msg := fmt.Sprintf("you are logged with %s", form.Username)
			c.JSON(http.StatusOK, gin.H{"status": msg})
		} else {
			c.JSON(401, gin.H{"status": "unauthorized"})
		}
	} else {
		c.JSON(401, gin.H{"status": "invalid infomation"}) //unable to bind
	}
}

type LoginFrom struct {
	Username string `form:"username" binding:"required"`
	Password string `form:"password" binding:"required"`
}
