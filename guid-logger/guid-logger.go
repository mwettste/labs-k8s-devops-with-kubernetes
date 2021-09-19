package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/google/uuid"
)

func main() {
	guid := uuid.New()

	ticker := time.NewTicker(3000 * time.Millisecond)
	done := make(chan bool)
	quitChan := make(chan os.Signal, 1)
	signal.Notify(quitChan, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		for {
			select {
			case <-done:
				return
			case <-ticker.C:
				logGuidAndTime(guid)
			}
		}
	}()

	logGuidAndTime(guid)

	<-quitChan
	ticker.Stop()
	done <- true
	fmt.Println("program ending...")
}

func logGuidAndTime(guid uuid.UUID) {
	nowStr := time.Now().Format(time.RFC3339)
	fmt.Printf("%s - %s\n", nowStr, guid.String())
}
