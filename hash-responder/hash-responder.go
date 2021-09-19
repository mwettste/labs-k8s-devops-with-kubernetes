package main

import (
	"crypto/md5"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Printf("Server started, listening on %s\n", os.Getenv("PORT"))
	r := gin.Default()

	startTimeString := time.Now().String()
	startHash := md5.Sum([]byte(startTimeString))

	r.GET("/", func(c *gin.Context) {
		nowTimeString := time.Now().String()
		nowHash := md5.Sum([]byte(nowTimeString))
		response := fmt.Sprintf("AppHash: %x - RequestHash: %x", startHash, nowHash)
		c.JSON(http.StatusOK, gin.H{
			"message": response,
		})
		fmt.Println(response)
	})
	r.Run()

	fmt.Println("server shutting down...")
}
