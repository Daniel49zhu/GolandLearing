package main

import "fmt"

func main() {
	messages := make(chan string)
	signals := make(chan bool)

	//Here’s a non-blocking receive.
	// If a value is available on messages
	// then select will take the <-messages case with that value.
	// If not it will immediately take the default case.
	select {
	case msg := <-messages:
		fmt.Println("received message " + msg)
	default:
		fmt.Println("no message received")
	}
	//A non-blocking send works similarly.
	// Here msg cannot be sent to the messages channel,
	// because the channel has no buffer and there is no receiver.
	msg := "hi"
	select {
	case messages <- msg:
		fmt.Println("received message", msg)
	default:
		fmt.Println("no message received")
	}
	//We can use multiple cases above
	// the default clause to implement a multi-way non-blocking select.
	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	case sig := <-signals:
		fmt.Println("received signal", sig)
	default:
		fmt.Println("no activity")
	}
}
