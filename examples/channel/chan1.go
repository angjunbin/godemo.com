package channel

import (
	"fmt"
	"sync"
)

// 优雅的关闭channel 写:读=> 1:N
func CloseChan() {
	// channel 初始化
	c := make(chan int, 10)
	// 用来 recevivers 同步事件的
	wg := sync.WaitGroup{}

	// sender（写端）
	go func() {
		// 入队
		c <- 1

		// 满足某些情况，则 close channel
		close(c)
	}()

	// receivers （读端）
	for i := 0; i < 10; i++ {
		wg.Add(1)
		fmt.Println("循环-", i)
		go func() {
			defer wg.Done()
			// ... 处理 channel 里的数据
			for v := range c {
				_ = v
				fmt.Println("循环的数据：", v)
			}
		}()
	}
	// 等待所有的 receivers 完成；
	wg.Wait()
}
