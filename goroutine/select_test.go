package goroutine

import (
	"fmt"
	"testing"
)

func TestSelect(t *testing.T) {
	ch := make(chan int, 1)
	for i := 0; i < 10; i++ {
		select {	//select会随机选择一项执行
		case x := <-ch:
			fmt.Println("已读出:", x)
		case ch <- i:
			fmt.Println("已写入:", i)

		}
	}
}