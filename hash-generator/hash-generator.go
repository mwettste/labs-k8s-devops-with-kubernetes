package main

import (
	"crypto/md5"
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(3000 * time.Millisecond)
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				return
			case <-ticker.C:
				printTimeNowHash()
			}
		}
	}()

	fmt.Scanln()
	ticker.Stop()
	done <- true
	fmt.Println("program ending...")
}

func printTimeNowHash() {
	nowStr := time.Now().String()
	hmd5 := md5.Sum([]byte(nowStr))
	fmt.Printf("%x\n", hmd5)
}
