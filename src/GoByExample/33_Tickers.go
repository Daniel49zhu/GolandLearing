package main

import (
	"fmt"
	"time"
)

// Timers are for when you want to do something once in the future
//tickers are for when you want to do something repeatedly at regular intervals.
func main() {
	ticker := time.NewTicker(500 * time.Millisecond)
	go func() {
		for t := range ticker.C {
			fmt.Println("Tick at", t)
		}
	}()
	//Timer是在将来要执行一次的事件
	//Ticker是将来要重复执行的事件
	time.Sleep(1600 * time.Millisecond)
	ticker.Stop()
	fmt.Println("Ticker stopped")
}
