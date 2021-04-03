package goroutine

import (
	"fmt"
	"sync"
	"testing"
)

var x int64
var wg sync.WaitGroup
var lock sync.Mutex

func add(no int) {
	for i := 0; i < 10; i++ {
		lock.Lock() // 加锁
		x = x + 1
		fmt.Println(no,":",x)
		lock.Unlock() // 解锁
	}
	wg.Done()
}

func TestLock(t *testing.T) {
	wg.Add(2)
	go add(1)
	go add(2)
	wg.Wait()
	//fmt.Println(x)
}
