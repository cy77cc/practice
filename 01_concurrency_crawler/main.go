package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
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

func main() {
	// 1. 构造 100 个模拟 URL
	urls := make([]string, 100)
	for i := 0; i < 100; i++ {
		urls[i] = fmt.Sprintf("http://example.com/%d", i)
	}

	// 2. 设置超时 Context (3秒)
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	fmt.Println("开始处理任务...")
	startTime := time.Now()

	// Jobs Channel 和 Results Channel
	jobs := make(chan string, 100) // 缓冲设大一点避免发送阻塞
	res := make(chan Result, 100)

	var wg sync.WaitGroup

	// 启动 5 个 Worker
	for i := 0; i < 5; i++ {
		wg.Go(func() {
			for {
				select {
				case <-ctx.Done():
					return
				case url, ok := <-jobs:
					if !ok {
						return // Channel 已关闭且无数据
					}
					status, duration, err := mockFetch(ctx, url)
					res <- Result{
						URL:      url,
						Status:   status,
						Duration: duration,
						Err:      err,
					}
				}
			}
		})
	}

	// 发送任务到 Jobs Channel
	go func() {
		for _, url := range urls {
			select {
			case <-ctx.Done():
				return // 超时提前退出
			case jobs <- url:
			}
		}
		close(jobs) // 发送完毕关闭 Channel
	}()

	// 监控任务完成状态，完成后关闭结果 Channel
	go func() {
		wg.Wait()
		close(res)
	}()

	// 收集结果并统计
	var successCount, failCount int
	var totalDuration time.Duration

loop:
	for {
		select {
		case <-ctx.Done():
			fmt.Println("\n任务超时强制结束！")
			break loop
		case r, ok := <-res:
			if !ok {
				// res Channel 被关闭，说明所有任务已完成
				break loop
			}
			// 打印单个结果
			// fmt.Printf("完成: %s Status: %d\n", r.URL, r.Status)

			// 统计
			if r.Status == 200 {
				successCount++
			} else {
				failCount++
			}
			totalDuration += r.Duration
		}
	}

	// 打印最终统计
	fmt.Println("\n=== 统计信息 ===")
	fmt.Printf("成功: %d\n", successCount)
	fmt.Printf("失败: %d\n", failCount)
	fmt.Printf("总任务耗时(累加): %v\n", totalDuration)
	fmt.Printf("程序运行总耗时: %v\n", time.Since(startTime))

}
