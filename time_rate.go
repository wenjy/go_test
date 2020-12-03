package main

import (
	"context"
	"fmt"
	"time"

	"golang.org/x/time/rate"
)

func main() {
	// 初始化捅是满的
	limiter := rate.NewLimiter(1, 10)

	// 拿出10个
	if r := limiter.AllowN(time.Now(), 10); r {
		fmt.Println("AllowN", r)
	}
	parent := context.Background()

	// 限时3s 产生3个
	ctx, cancelFunc := context.WithTimeout(parent, time.Second*3)

	// 取4个不成功，调用取消
	err := limiter.WaitN(ctx, 4)
	if err != nil {
		cancelFunc()
		fmt.Println("cannel", err)
	}

	// 要准备拿3个，目前数量为0
	r := limiter.ReserveN(time.Now(), 3)
	if !r.OK() {
		return
	}

	// 不想等待可以调用 r.Cancel() 取消
	if r.Delay() > 0 {
		fmt.Println("Sleep", r.Delay())
		time.Sleep(r.Delay())
		fmt.Println("ReserveN act")
	}

	// 重新设置速度
	limiter.SetLimit(2)

	// 拿出3个，没有
	if r := limiter.AllowN(time.Now(), 3); r {
		fmt.Println("AllowN2", r)
	}

	// 要准备拿4个，目前数量为0，要sleep 2s
	r = limiter.ReserveN(time.Now(), 4)
	if !r.OK() {
		return
	}

	// 不想等待可以调用 r.Cancel() 取消
	if r.Delay() > 0 {
		fmt.Println("Sleep", r.Delay())
		time.Sleep(r.Delay())
		fmt.Println("ReserveN act")
	}
}
