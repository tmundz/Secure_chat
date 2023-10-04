package main

import (
	"http_server/database"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	database.ConnectDb()
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "hello world"})
	})
	r.Run()
}
