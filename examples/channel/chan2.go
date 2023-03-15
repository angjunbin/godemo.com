package channel

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func CloseChanMulSend() {
	// channel 初始化
	c := make(chan int, 10)
	// 用来 recevivers 同步事件的
	wg := sync.WaitGroup{}
	// 上下文
	ctx, cancel := context.WithCancel(context.TODO())

	// 专门关闭的协程
	go func() {
		time.Sleep(2 * time.Second)
		cancel()

		// ... 某种条件下关闭channel
		fmt.Println("2 分钟超时之后，单独协程执行 close(channel) 操作")
		close(c)
	}()

	count := 10

	// senders
	for i := 0; i < count; i++ {
		go func(ctx context.Context, id int) {
			select {
			case <-ctx.Done():
				return
			case c <- id: // 入队
				fmt.Println("入队：", id)
			}
		}(ctx, i)
	}

	// receivers
	for i := 0; i < count; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			// 处理chan数据
			for v := range c {
				_ = v
				fmt.Println("v=", v)
			}
		}()
	}

	wg.Wait()
}
