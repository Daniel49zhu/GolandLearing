package main

import "fmt"

// Closing a channel indicates that no more values will be sent on it.
func main() {
	jobs := make(chan int, 5)
	done := make(chan bool)

	go func() {
		for {
			j, more := <-jobs
			if more {
				fmt.Println("received job", j)
			} else {
				fmt.Println("received all jobs")
				done <- true
				return
			}
		}
	}()

	for j := 1; j <= 3; j++ {
		jobs <- j
		fmt.Println("send job", j)
	}
	// the more value will be false if jobs has been closed and all values in the channel have already been received.
	close(jobs)
	fmt.Println("send all jobs")

	<-done
}
