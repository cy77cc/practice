# 题目 02: 多路日志库设计

## 背景
在项目中，我们通常需要将日志输出到不同的地方：控制台、文件、或者远程日志收集服务（如 Kafka, Elasticsearch）。Go 语言的接口（Interface）机制非常适合实现这种可插拔的设计。

## 需求
设计并实现一个日志库 `Logger`，要求如下：
1.  **接口定义**：定义一个 `LogWriter` 接口，包含 `Write(level string, msg string)` 方法。
2.  **实现类**：
    *   `ConsoleWriter`：将日志格式化输出到标准输出（os.Stdout）。
    *   `FileWriter`：将日志写入到指定的文件中。
3.  **核心 Logger**：
    *   支持注册多个 `LogWriter`。
    *   提供 `Info(msg string)`, `Error(msg string)` 方法。
    *   当调用 `Logger.Info` 或 `Logger.Error` 时，所有注册的 Writer 都应该收到日志并执行写入。
4.  **格式要求**：日志内容应包含时间戳、级别和消息内容。例如：`[2023-10-01 12:00:00] [INFO] This is a message`。

## 考察点
*   Interface 的定义与实现（多态）
*   结构体组合与切片操作
*   文件 I/O 操作 (`os.OpenFile`, `os.O_APPEND`)
*   依赖注入思想

## 扩展（可选）
*   增加一个 `AsyncWriter`，使用 Channel 异步写入日志，不阻塞主业务流程。
