# 题目 03: HTTP 中间件机制

## 背景
中间件（Middleware）是 Web 框架的核心组件，用于处理横切关注点，如日志记录、鉴权、Panic 恢复、请求耗时统计等。

## 需求
使用 Go 标准库 `net/http` 实现一个带有中间件机制的 HTTP Server。

1.  **实现中间件类型**：
    定义中间件函数类型 `type Middleware func(http.Handler) http.Handler`。

2.  **编写三个具体的中间件**：
    *   `LoggingMiddleware`: 记录每个请求的方法、URL 和处理耗时。
    *   `AuthMiddleware`: 检查请求头 `Authorization`。如果值不是 "secret-token"，则返回 401 Unauthorized，拦截请求。
    *   `RecoveryMiddleware`: 捕获处理过程中可能发生的 Panic，防止服务崩溃，并返回 500 Internal Server Error。

3.  **业务处理 Handler**：
    *   编写一个 `HelloHandler`，返回 "Hello, World!"。
    *   编写一个 `PanicHandler`，手动触发 panic，用于测试 Recovery 中间件。

4.  **链式调用**：
    *   编写一个辅助函数 `Chain(h http.Handler, m ...Middleware) http.Handler`，将多个中间件应用到 Handler 上。
    *   应用顺序建议：`Recovery` -> `Logging` -> `Auth` -> `Handler`。

## 考察点
*   `http.Handler` 接口理解
*   闭包（Closure）的使用
*   `defer` 和 `recover` 处理 Panic
*   HTTP 状态码控制

## 提示
*   `http.HandlerFunc` 可以将普通函数转换为 `http.Handler`。
*   中间件通常包装原始 Handler，在调用 `next.ServeHTTP` 之前或之后执行逻辑。
