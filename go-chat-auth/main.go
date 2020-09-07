package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/test", Test)
	err := router.Run(":8090")
	if err != nil {
		log.Fatal(err)
	}
}

func Test(c *gin.Context) {
	c.String(http.StatusOK, "wo ai go-chat-auth")
}
