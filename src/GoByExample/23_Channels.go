package main

import "fmt"

// Channels are the pipes that connect concurrent goroutines.

func main() {
	message := make(chan string)

	//启动一个goroutine
	go func() {
		//Send a value into a channel using the channel <- syntax.
		//消息进入通道
		message <- "ping"
	}()
	//消息出通道 传给message
	msg := <-message
	fmt.Println(msg)
}
