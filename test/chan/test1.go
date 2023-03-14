package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)
	go func() {
		fmt.Println("向channel 发送数据：")
		ch <- "string1"
		time.Sleep(time.Second * 1)
	}()

	s1 := <-ch
	time.Sleep(time.Second * 2)
	fmt.Println("接受的数据：", s1)

}
