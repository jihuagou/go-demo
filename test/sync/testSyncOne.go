package sync

import (
	"fmt"
	"sync"
	"time"
)

var (
	once sync.Once
)

func SyncOnceDemo()  {
	// 只执行一个T1函数，输出T1
	once.Do(T1)
	once.Do(T2)
	time.Sleep(time.Second)
}
func T1()  {
	fmt.Println("T1")
}

func T2()  {
	fmt.Println("T1")
}
