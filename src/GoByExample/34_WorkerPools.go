package main

import (
	"fmt"
	"time"
)

func worker2(id int, jobs <-chan int, result chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "started job", j)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished job", j)
		result <- j * 2
	}
}

func main() {
	jobs := make(chan int, 100)
	results := make(chan int, 100)
	//启动三个goroutine，因为job中暂时没有数据所有都阻塞
	for w := 1; w <= 3; w++ {
		go worker2(w, jobs, results)
	}
	//向jobs中依次传入5个数，三个线程开始工作
	for j := 1; j <= 5; j++ {
		jobs <- j
	}
	close(jobs)
	for a := 1; a <= 5; a++ {
		<-results
	}
}
