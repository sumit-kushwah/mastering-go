package concurrency

// https://www.freecodecamp.org/news/concurrent-programming-in-go/
// https://www.rudderstack.com/blog/implementing-graceful-shutdown-in-go/

import (
	osmodule "github.com/sumit-kushwah/mastering-go-concurrency/os-module"
)

func writeToFileForOneMinute() {
	osmodule.WriteToFile("/tmp/test.txt", "My name is sumit\n")
}

// https://stackoverflow.com/questions/72553044/what-happens-to-unfinished-goroutines-when-the-main-parent-goroutine-exits-or-re
func MainDoesNotWait() {
	go writeToFileForOneMinute()
	// use below line if you want to wait for the go routine to finish
	// otherwise the program will exit before the go routine finishes
	// and all go routines will abruptly stop
	// time.Sleep(1 * time.Second)
}
