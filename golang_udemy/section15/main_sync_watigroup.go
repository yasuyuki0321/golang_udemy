package main

import (
	"fmt"
	"sync"
	"time"
)

// sync

func main() {
	wg := new(sync.WaitGroup)
	wg.Add(3)

	go func() {
		for i := 0; i < 100; i++ {
			fmt.Println("1st Goroutine")
			time.Sleep(time.Millisecond)
		}
		wg.Done()
	}()
	go func() {
		for i := 0; i < 100; i++ {
			fmt.Println("2st Goroutine")
			time.Sleep(time.Millisecond)
		}
		wg.Done()
	}()
	go func() {
		for i := 0; i < 100; i++ {
			fmt.Println("3st Goroutine")
			time.Sleep(time.Millisecond)
		}
		wg.Done()
	}()

	// Doneが3つになるまでまつ
	wg.Wait()
	// for {
	// }
}
