package concurrent

import (
	"context"
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestContextDone(t *testing.T) {
	fmt.Println("处理请求当中")

	ctx, cancelFunc := context.WithCancel(context.Background())

	go calculatePos(ctx)

	time.Sleep(time.Second * 10)

	fmt.Println("客户端发送通知，用户已经退出骑手距离用户界面")
	cancelFunc()
}

//处理计算地理位置
func calculatePos(ctx context.Context) {
	for {
		fmt.Println("发送实时计算位置")
		select {
		case <-ctx.Done():
			return
		case <-time.After(time.Second):
		}
	}
}

func gen(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

func sq(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}

//合并多个chan
func merge(cs ...<-chan int) <-chan int {
	result := make(chan int)

	wg := &sync.WaitGroup{}
	mergeChannel := func(in <-chan int) {
		defer wg.Done()
		for val := range in {
			result <- val
		}
	}

	wg.Add(len(cs))
	for _, c := range cs {
		go mergeChannel(c)
	}

	go func() {
		wg.Wait()
		close(result)
	}()
	return result
}
