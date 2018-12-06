package main

import (
	"fmt"
	"time"
)

func worker(done chan bool) {
	fmt.Println("working....")
	time.Sleep(time.Second)
	fmt.Println("Done")
	//传输一个信号来告诉主线程执行完毕了
	done <- true
}

func main() {
	done := make(chan bool, 1)
	go worker(done)
	//主线程在此阻塞 直到从done通道中接收到参数才会推出
	//如果注释掉这句，worker线程还没开启主线程就已经结束退出了
	<-done
}
