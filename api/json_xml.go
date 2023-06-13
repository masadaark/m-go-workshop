package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SomeJSON(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "hey"})
}

func MoreJSON(c *gin.Context) {
	var msg struct {
		Name    string `json:"user"`
		Message string `json:"message"`
		Id      int    `json:"id"`
	}
	msg.Name = "Mark"
	msg.Message = "Hello"
	msg.Id = 33163
	c.JSON(http.StatusOK, msg)
}

func SomeXML(c *gin.Context) {
	c.XML(http.StatusOK, gin.H{"message": "Hello"})
}

func SomeYAML(c *gin.Context) {
	c.YAML(http.StatusOK, gin.H{"message": "Hello"})
}
