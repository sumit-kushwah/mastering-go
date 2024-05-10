package concurrency

import (
	"fmt"
	"runtime"
	"strconv"
)

func printName() {
	i := 0
	for {
		i++
		fmt.Println("sumit")
		if i == 10 {
			break
		}
	}
}

func printName2() {
	fmt.Println("sumit 2")
}

func MainForever() {
	forever := make(chan bool)
	printName2()
	fmt.Println("Current Goroutine Count: " + strconv.Itoa(runtime.NumGoroutine()))
	go func(ch chan bool) {
		i := 0
		for {
			i++
			fmt.Println("sumit")
			if i == 10 {
				break
			}
		}
		close(forever)
		fmt.Println("sumit 2")
		fmt.Println("Current Goroutine Count: " + strconv.Itoa(runtime.NumGoroutine()))
	}(forever)
	// time.Sleep(time.Second * 2)
	fmt.Println("Current Goroutine Count: " + strconv.Itoa(runtime.NumGoroutine()))
	fmt.Println(" [*] Waiting for queue data. To exit press CTRL+C")
	<-forever
	fmt.Println(" [*] Exiting")
}
