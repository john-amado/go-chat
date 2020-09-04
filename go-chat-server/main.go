package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("hello world")

	InitRouter()
}

func InitRouter() {
	router := gin.Default()
	router.GET("/test", Test)
	err := router.Run(":9090")
	if err != nil {
		log.Fatal(err)
	}
}

func Test(c *gin.Context) {
	c.String(http.StatusOK, "wo ai ni")
}
