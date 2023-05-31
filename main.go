package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

var logInCount int = 0

func main() {
	r := gin.Default() //create router variable

	//control log output coloring
	gin.DisableConsoleColor()
	gin.ForceConsoleColor()
	//logger
	runningDir, _ := os.Getwd()
	errorlogfile, _ := os.OpenFile(fmt.Sprintf("%s/gin_error.log", runningDir), os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0600) //createdFile|ifAlreadyHasFile(FlagA+)|WriteOnly
	accesslogfile, _ := os.OpenFile(fmt.Sprintf("%s/access.log", runningDir), os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0600)

	//ผูกไฟล์ error กับ logger
	gin.DefaultErrorWriter = errorlogfile
	gin.DefaultWriter = accesslogfile
	// r.Use(gin.Logger()) //standardLogger
	/*r.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s\"%s\" %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC1123),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))*/
	r.Use(gin.LoggerWithWriter(gin.DefaultWriter, "/profile")) //hide this route in access log

	r.GET("/", func(c *gin.Context) {
		c.Data(http.StatusOK, "text/html; charset=uft-8", []byte("Root")) //(httpStatus,contentType,response)
	})

	r.GET("/test", func(c *gin.Context) {
		c.Data(200, "text/html; charset=uft-8", []byte("test"))
	})
	//query param 5/login?username=hua&password=jai
	r.GET("/login", func(c *gin.Context) {
		logInCount += 1
		accesslogfile.WriteString(fmt.Sprintf("login count : %d", logInCount))
		username, password := c.Query("username"), c.Query("password")
		c.JSON(http.StatusOK, gin.H{"result": "ok", "username": username, "password": password})
	})

	r.GET("/error", func(c *gin.Context) {
		errorlogfile.WriteString(fmt.Sprintf("\nerror : %s\n", "เว้นบรรทัด"))
		c.JSON(888, gin.H{"reason": "หล่อเกินไป"})
	})

	r.POST("/login", postLogin)

	r.GET("/fullName/:firstName/:lastName", returnFullName)

	//Set lower Moemory limit for multipart forms (default is 32 MiB)
	r.MaxMultipartMemory = 8 << 20 //8 MiB

	r.POST("/upload", upLoadFile)

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

func upLoadFile(c *gin.Context) {
	// single File
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "No file uploaded"})
		return
	}
	extension := filepath.Ext(file.Filename)

	runningDir, _ := os.Getwd()

	err = c.SaveUploadedFile(file, fmt.Sprintf("%s/upload/%s%s", runningDir, file.Filename, extension))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "Failed to save file"})
		return
	}

	c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded", file.Filename))
}
