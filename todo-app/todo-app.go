package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Printf("Server started, listening on %s\n", os.Getenv("PORT"))
	r := gin.Default()
	r.GET("ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.Run()
}
