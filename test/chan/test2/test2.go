package main

import (
	"fmt"
	"time"
)

func producter(ch chan<- int, n int) {
	for i := 0; i < n; i++ {
		ch <- i
		fmt.Println("Send data:", i)
	}
}

func consumer(ch <-chan int, n int) {
	for i := 0; i < n; i++ {
		data := <-ch
		fmt.Println("接收的数据：", data)
	}
}

// 生产者消费者的例子
func main() {
	ch := make(chan int)
	go producter(ch, 10)
	go consumer(ch, 10)
	time.Sleep(time.Second * 1)
}
