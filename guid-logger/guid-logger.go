package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func main() {
	guid := uuid.New()
	r := gin.Default()

	ticker := time.NewTicker(3000 * time.Millisecond)
	done := make(chan bool)
	quitChan := make(chan os.Signal, 1)
	signal.Notify(quitChan, syscall.SIGINT, syscall.SIGTERM)
	var logString string

	go func() {
		for {
			select {
			case <-done:
				return
			case <-ticker.C:
				logString = logGuidAndTime(guid)
			}
		}
	}()

	logGuidAndTime(guid)

	r.GET("log", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": logString,
		})
	})
	r.Run()

	<-quitChan
	ticker.Stop()
	done <- true
	fmt.Println("program ending...")
}

func logGuidAndTime(guid uuid.UUID) string {
	nowStr := time.Now().Format(time.RFC3339)
	logString := fmt.Sprintf("%s - %s", nowStr, guid.String())
	fmt.Println(logString)
	return logString
}
