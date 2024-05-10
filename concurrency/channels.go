package concurrency

import (
	"fmt"
	"strconv"
	"time"
)

func channelHelper(ch chan string) {
	// time.Sleep(time.Second * 2)
	ch <- "Hello"
	ch <- "Sumit"
	ch <- "Kushwah"
	// This code is blocked until the channel value is consumed
	close(ch)
	fmt.Println("Channel Closed")
}

func channelHelperWithBuffer(ch chan string) {
	for i := 0; i < 20; i++ {
		ch <- "Hello " + strconv.Itoa(i)
		// fmt.Println("Sent " + strconv.Itoa(i))
	}
	close(ch) // important to close the channel
}

func MainBufferedChannels() {
	var ch chan string = make(chan string, 5)
	go channelHelperWithBuffer(ch)
	// This for loop will give issue if we don't close the channel
	// after sending all the values
	for value := range ch {
		fmt.Println(value)
		time.Sleep(time.Millisecond * 100)
	}
	// This code is blocked until the channel receives a value
	fmt.Println("my name is sumit")
}

func MainUnBufferedChannels() {
	var ch chan string = make(chan string)
	go channelHelper(ch)
	// time.Sleep(time.Second * 3)
	for v := range ch {
		fmt.Println(v)
		time.Sleep(time.Second * 2)
	}
	// fmt.Println(<-ch)
	// This code is blocked until the channel receives a value
	fmt.Println("my name is sumit")
}
