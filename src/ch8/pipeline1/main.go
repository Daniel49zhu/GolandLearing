package main

import (
	"fmt"
	"time"
)

func main() {
	naturals := make(chan int) //建立一个通道
	squares := make(chan int)  //建立一个通道

	go func() {
		for x := 0; ; x++ {
			//向natutals通道传递一个自然数
			naturals <- x
		}
	}()

	go func() {
		for {
			//从naturals通道接收一个自然数
			x := <-naturals
			//平方之后传递给squares通道
			squares <- x * x
		}
	}()

	for {
		fmt.Println(<-squares)
		time.Sleep(2 * time.Second)
	}
}
