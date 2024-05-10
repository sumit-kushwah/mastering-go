package concurrency

import (
	"fmt"
	"sync"
	"time"
)

func routine1(wg *sync.WaitGroup) {
	fmt.Println("Routine 1")
	wg.Done()
}
func routine2(wg *sync.WaitGroup) {
	fmt.Println("Routine 2")
	wg.Done()
}

func MainWaitGroupUseCase() {
	wg := new(sync.WaitGroup)
	wg.Add(2)
	go routine1(wg)
	go routine2(wg)
	wg.Wait()
}

// Running Max 10 go routines at a time
func MainMax10GoRoutines() {
	maxRoutines := 10
	wg := new(sync.WaitGroup)
	for i := 1; i <= 100; i++ {
		wg.Add(1)
		go func(i int) {
			fmt.Printf("Routine %d\n", i)
			time.Sleep(5 * time.Second)
			wg.Done()
		}(i)
		if i%maxRoutines == 0 {
			wg.Wait()
		}
	}
	wg.Wait()
}
