package main

import (
	"fmt"
)

// TODO: 定义 LogWriter 接口
type LogWriter interface {
	// Write(level string, msg string) error
}

// TODO: 实现 ConsoleWriter
type ConsoleWriter struct{}

// TODO: 实现 FileWriter
type FileWriter struct {
	FilePath string
}

// Logger 核心结构体
type Logger struct {
	writers []LogWriter
}

// Register 注册 Writer
func (l *Logger) Register(w LogWriter) {
	l.writers = append(l.writers, w)
}

// Info 记录 Info 级别日志
func (l *Logger) Info(msg string) {
	// TODO: 遍历 writers 调用 Write
}

// Error 记录 Error 级别日志
func (l *Logger) Error(msg string) {
	// TODO: 遍历 writers 调用 Write
}

func main() {
	logger := &Logger{}

	// 1. 注册控制台输出
	// logger.Register(&ConsoleWriter{})

	// 2. 注册文件输出
	// logger.Register(&FileWriter{FilePath: "app.log"})

	// 3. 打印日志
	logger.Info("应用启动成功")
	logger.Error("连接数据库失败")

	fmt.Println("日志演示结束，请检查控制台和 app.log 文件")
}
