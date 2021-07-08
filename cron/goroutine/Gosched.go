package goroutine

import (
	"fmt"
	"runtime"
	"time"
)

func RunGosched()  {
	go func(s string) {
		for i := 0; i < 2; i++ {
			fmt.Println(s)
		}
	}("world")

	// 主协程
	for i := 0; i < 2; i++ {
		// 切一下，再次分配任务
		runtime.Gosched()
		fmt.Println("hello")
	}
}

// runtime.GOMAXPROCS() 确定需要多少个OS线程来同时执行Go代码。默认值是机器上的CPU核心数。
func RunGomaxProcs() {
	runtime.GOMAXPROCS(10)
	startTime := time.Now()
	//fmt.Println(startTime.UnixNano())
	go func() {
		for i := 0; i < 1000; i++ {
			//fmt.Println("A:", i)
		}
	}()
	go func() {
		for i := 0; i < 1000; i++ {
			//fmt.Println("B:", i)
		}
	}()
	endTime := time.Now()
	fmt.Println(endTime.UnixNano() - startTime.UnixNano())
	time.Sleep(time.Second)
}