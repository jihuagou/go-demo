package main

import (
	"fmt"
	"time"
)

func main()  {
	testCacheTime()
}

const NUM = 9999

// 测试程序中利用缓存的局部性原理，大大提高程序的运行效率
// 局部性分为：时间局部性和空间局部性。如果一个内存位置被重复的引用，那就是有了时间局部性。
//  如果一个内存位置被引用了，很快这个位置的附件位置也被引用了，这既是空间局部性
func testCacheTime() {

	// 初始化数据
	data := [NUM][NUM]int{}
	for i := 0; i < NUM; i++ {
		for j := 0; j < NUM; j++ {
			data[i][j] = 1
		}
	}
	sum := 0

	beginUnix := time.Now()

	for i := 0; i < NUM; i++ {
		for j := 0; j < NUM; j++ {
			sum += data[i][j]
		}
	}

	fmt.Println(sum, time.Now().UnixNano() - beginUnix.UnixNano())

	sum = 0
	beginUnix = time.Now()
	for i := 0; i < NUM; i++ {
		for j := 0; j < NUM; j++ {
			sum += data[j][i]
		}
	}

	fmt.Println(sum, time.Now().UnixNano() - beginUnix.UnixNano())
}
