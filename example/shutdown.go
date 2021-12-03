package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"

	"github.com/munir131/grace"
)

func main() {
	grace.Init(2 * time.Second)
	go func() {
		for index := 0; index < 100; index++ {
			go func() {
				time.Sleep(time.Duration(rand.Int31n(5)) * time.Second)
			}()
		}
	}()
	go func() {
		time.Sleep(1 * time.Second)
		grace.SetTrue()
	}()

	fmt.Println(runtime.NumGoroutine(), " are currently running")
	check := grace.CheckGraceSig()

	fmt.Println("signal not received yet", check)
	time.Sleep(2 * time.Second)

	check = grace.CheckGraceSig()
	fmt.Println("signal received ", check)
	// Time to shutdown
	for !grace.PeriodOver() && runtime.NumGoroutine() != 1 {
		fmt.Printf("Can't shutdown as %d goroutines are still running\n", runtime.NumGoroutine())
		time.Sleep(1 * time.Second)
	}
}
