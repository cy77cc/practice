# 题目 01: 并发URL处理器 (Worker Pool 模型)

## 背景
在后端开发中，经常需要处理大量的并发任务，例如批量发送HTTP请求、批量处理数据等。控制并发数量、收集处理结果、优雅退出是必备技能。

## 需求
实现一个并发URL处理器，要求如下：
1.  **输入**：一个包含 100 个 URL 的切片（模拟数据，如 `http://example.com/1` ... `http://example.com/100`）。
2.  **处理**：模拟对每个 URL 进行 HTTP 请求（可以用 `time.Sleep` 模拟耗时 100ms ~ 500ms），并返回状态码（模拟 200 或 500）。
3.  **并发控制**：限制同时工作的 Worker 数量为 5 个。不能开启 100 个 Goroutine。
4.  **结果收集**：收集所有请求的结果（URL, 状态码, 耗时），并在所有任务完成后打印统计信息（成功数、失败数、总耗时）。
5.  **超时控制**：整个任务如果超过 3 秒未完成，强制结束并打印已完成的结果。

## 考察点
*   Goroutine 与 Channel 的使用
*   Worker Pool 模式的实现
*   `sync.WaitGroup` 
*   `context` 包进行超时控制
*   `sync.Mutex` 或 Channel 进行并发安全的数据收集

## 提示
*   创建一个 Jobs Channel 传递 URL。
*   创建一个 Results Channel 传递处理结果。
*   启动固定数量的 Worker Goroutine 消费 Jobs Channel。
