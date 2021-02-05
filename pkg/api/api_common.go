package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"data":   data,
		"status": "success",
	})
}

func Fail(c *gin.Context, message string, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"data":    data,
		"status":  "fail",
		"message": message,
	})
}

func Errorf(format string, val ...interface{}) {
	fmt.Printf("%s [ERROR] ", time.Now().Format("2006-01-02 15:04:05"))
	fmt.Printf(format, val...)
	fmt.Print("\n")
}

func Infof(format string, val ...interface{}) {
	fmt.Printf("%s [INFO] ", time.Now().Format("2006-01-02 15:04:05"))
	fmt.Printf(format, val...)
	fmt.Print("\n")
}
