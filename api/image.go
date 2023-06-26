package api

import (
	"github.com/gin-gonic/gin"
)

func Jung(c *gin.Context) {
	c.Header("Content-Description", "Simulation File Download")
	c.Header("Content-Transfer-Encoding", "binary")
	c.Header("Content-Disposition", "attachment; filename=jungyard.png")
	c.Header("Content-Type", "application/octet-stream")
	c.File("file/jungyard.png")
}
