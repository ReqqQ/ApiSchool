package main

import (
	"github.com/gin-gonic/gin"
)

const RESPONSE_GET string = "/ping"

func main() {
	api := gin.Default()
	api.GET(RESPONSE_GET, func(response *gin.Context) {
		response.JSON(200, gin.H{
			"message": "pong",
		})
	})
	api.Run()
}
