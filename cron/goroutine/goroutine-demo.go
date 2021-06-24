package goroutine

import (
	"fmt"
	"time"
)

func Start()  {
	demo1()
	demo2()
}

// 定期轮询channel
func demo2()  {
	done := make(chan struct{})
	ch := make(chan string)

	// 通过done来控制ch结束
	// 3秒后结束ch，从而结束goroutine
	go func() {
		time.Sleep(3 * time.Second)
		done <- struct{}{}
	}()

	go func() {
		for true {
			select {
			case ch <- "this is a test":
			case <- done:
				close(ch)
				return
			}
			time.Sleep(100 * time.Millisecond)
		}
	}()

	for i := range ch {
		fmt.Println("ch接受到的值", i)
	}

	fmt.Println("Over!")

}


// 关闭channel例子
// 1. 借助channel的close机制来完成对goroutine的精确控制
//       如果程序逻辑比较简单，可以不调用close方法，因为goroutine会自然结束，也就不需要手动关闭
func demo1()  {
	ch := make(chan string, 0)

	go func() {
		for true {
			v, ok := <-ch
			if !ok {
				fmt.Println("goroutine closed")
				return
			}

			fmt.Println(v)
		}


	}()

	ch <- "a"
	ch <- "b"
	close(ch)
	time.Sleep(time.Second)
}
