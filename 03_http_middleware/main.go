package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

// Middleware 定义中间件类型
type Middleware func(http.Handler) http.Handler

// LoggingMiddleware 日志中间件
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		log.Printf("Started %s %s", r.Method, r.URL.Path)
		
		next.ServeHTTP(w, r)
		
		log.Printf("Completed %s in %v", r.URL.Path, time.Since(start))
	})
}

// TODO: 实现 AuthMiddleware
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 检查 Header "Authorization" == "secret-token"
		// 失败则 w.WriteHeader(http.StatusUnauthorized) 并 return
		next.ServeHTTP(w, r)
	})
}

// TODO: 实现 RecoveryMiddleware
func RecoveryMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 使用 defer recover() 捕获 panic
		next.ServeHTTP(w, r)
	})
}

// 业务 Handler
func HelloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, World!")
}

func PanicHandler(w http.ResponseWriter, r *http.Request) {
	panic("Something went wrong!")
}

// Chain 辅助函数，将中间件应用到 Handler
func Chain(h http.Handler, middlewares ...Middleware) http.Handler {
	for i := len(middlewares) - 1; i >= 0; i-- {
		h = middlewares[i](h)
	}
	return h
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", HelloHandler)
	mux.HandleFunc("/panic", PanicHandler)

	// 应用中间件链
	// 注意顺序：Recovery 应该在最外层（最先进入，最后退出）以捕获所有 panic
	// Logging 其次
	// Auth 最后（最接近业务逻辑）
	
	// 这里目前只应用了 Logging，请补充其他中间件
	handler := Chain(mux, LoggingMiddleware)

	server := &http.Server{
		Addr:    ":8080",
		Handler: handler,
	}

	log.Println("Server starting on :8080...")
	log.Fatal(server.ListenAndServe())
}
