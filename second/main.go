package main

import (
	"fmt"
	"time"
)

func main() {
	sync := make(chan bool) // решение make(chan bool, 2)

	go func() {
		time.Sleep(time.Second * 3)
		fmt.Println("get normal signal")
		sync <- false
	}()
	ticker := time.NewTicker(time.Second)

	for {
		select {
		case <-ticker.C:
			fmt.Println("get interrupted signal")
			sync <- true
		case value := <-sync:
			fmt.Printf("finish %t", value)
			return
		}
	}
}
