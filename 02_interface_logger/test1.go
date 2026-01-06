package main

import (
	"fmt"
	"os"
	"time"
)

// LogWriter 接口
type LogWriter interface {
	Write(level string, msg string) error
}

// ---------------- ConsoleWriter ----------------

type ConsoleWriter struct{}

func (c *ConsoleWriter) Write(level string, msg string) error {
	// 直接写入 os.Stdout，避免频繁创建 bufio
	// 格式化时间
	ts := time.Now().Format("2006-01-02 15:04:05")
	_, err := fmt.Fprintf(os.Stdout, "ts=%s level=%s msg=%s\n", ts, level, msg)
	return err
}

// ---------------- FileWriter ----------------

type FileWriter struct {
	file *os.File // 持有文件句柄，而不是每次打开
}

// NewFileWriter 构造函数，负责打开文件
func NewFileWriter(filePath string) (*FileWriter, error) {
	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return nil, err
	}
	return &FileWriter{file: file}, nil
}

// Close 关闭文件资源
func (f *FileWriter) Close() error {
	if f.file != nil {
		return f.file.Close()
	}
	return nil
}

func (f *FileWriter) Write(level string, msg string) error {
	if f.file == nil {
		return fmt.Errorf("file is not open")
	}
	ts := time.Now().Format("2006-01-02 15:04:05")
	// 直接写入文件
	_, err := fmt.Fprintf(f.file, "ts=%s level=%s msg=%s\n", ts, level, msg)
	return err
}

// ---------------- Logger ----------------

type Logger struct {
	writers []LogWriter
}

func (l *Logger) Register(w LogWriter) {
	l.writers = append(l.writers, w)
}

func (l *Logger) Info(msg string) {
	for _, w := range l.writers {
		if err := w.Write("INFO", msg); err != nil {
			// 简单的错误兜底，输出到 stderr，不影响主流程
			fmt.Fprintf(os.Stderr, "Logger Error: %v\n", err)
		}
	}
}

func (l *Logger) Error(msg string) {
	for _, w := range l.writers {
		if err := w.Write("ERROR", msg); err != nil {
			fmt.Fprintf(os.Stderr, "Logger Error: %v\n", err)
		}
	}
}

// ---------------- Main ----------------

func main() {
	logger := &Logger{}

	// 1. 注册控制台输出
	logger.Register(&ConsoleWriter{})

	// 2. 注册文件输出 (使用构造函数)
	fileWriter, err := NewFileWriter("app.log")
	if err != nil {
		fmt.Printf("无法创建文件日志: %v\n", err)
		return
	}
	// 确保在 main 退出时关闭文件
	defer fileWriter.Close()

	logger.Register(fileWriter)

	// 3. 打印日志
	logger.Info("应用启动成功")
	logger.Error("连接数据库失败")

	fmt.Println("日志演示结束，请检查控制台和 app.log 文件")
}
