package main

import (
	"fmt"
	"time"
)

//We often want to execute Go code
// at some point in the future, or repeatedly at some interval.
func main() {
	//Timers represent a single event in the future. You tell the timer how long you want to wait,
	timer1 := time.NewTimer(2 * time.Second)

	<-timer1.C
	fmt.Println("Timer 1 expired")

	timer2 := time.NewTimer(time.Second)
	go func() {
		<-timer2.C
		fmt.Println("Timer 2 expired")
	}()
	// Stop() can cancel the timer before it expires.
	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("Timer 2 stopped")
	}
}
