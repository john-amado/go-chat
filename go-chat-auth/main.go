package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
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
