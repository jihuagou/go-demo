package main

import (
	"github.com/tmaio/go-gin-example/cron/goroutine"
)

// 定时任务入口
func main()  {
	//userCron.DoUserJob()
	goroutine.Start()
}
