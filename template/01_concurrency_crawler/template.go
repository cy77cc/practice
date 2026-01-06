package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

// Result 结构体用于保存处理结果
type Result struct {
	URL      string
	Status   int
	Duration time.Duration
	Err      error
}

// 模拟 HTTP 请求，返回状态码和耗时
func mockFetch(ctx context.Context, url string) (int, time.Duration, error) {
	// 模拟随机耗时
	duration := time.Duration(rand.Intn(400)+100) * time.Millisecond

	select {
	case <-ctx.Done():
		return 0, 0, ctx.Err()
	case <-time.After(duration):
		// 模拟随机状态码
		status := 200
		if rand.Intn(10) > 8 { // 10% 概率失败
			status = 500
		}
		return status, duration, nil
	}
}

func main1() {
	// 1. 构造 100 个模拟 URL
	urls := make([]string, 100)
	for i := 0; i < 100; i++ {
		urls[i] = fmt.Sprintf("http://example.com/%d", i)
	}

	// 2. 设置超时 Context (3秒)
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	// Jobs Channel 和 Results Channel

	// 启动 5 个 Worker

	// 发送任务到 Jobs Channel

	// 监控任务完成状态，完成后关闭结果 Channel

	// 收集结果并统计

	// 打印最终统计

}
