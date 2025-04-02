package backend

import "time"

func Start() {

	// 启动一些异步任务
	go func() {
		for {
			// 执行任务
			// ...
			// 等待一段时间
			<-time.After(time.Second * 10)
		}
	}()

}
