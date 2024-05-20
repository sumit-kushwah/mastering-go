package concurrency

import (
	"fmt"
	"math/rand"
	"time"
)

func printIndia(ch chan bool) {
	defer func() {
		<-ch
	}()
	fmt.Println("India " + time.Now().String())
	time.Sleep(1*time.Second + time.Duration(rand.Intn(500))*time.Millisecond)
}

func MainMax10Goroutine() {
	ch := make(chan bool, 5)
	for {
		ch <- true
		go printIndia(ch)
		return
	}
}
