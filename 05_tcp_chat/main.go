package main

import (
	"fmt"
	"log"
	"net"
)

// Client 代表一个聊天客户端
type Client struct {
	C    chan string // 用于发送消息给该客户端
	Name string
	Conn net.Conn
}

// 消息通道
var (
	entering = make(chan Client)
	leaving  = make(chan Client)
	messages = make(chan string) // 所有客户端发来的消息
)

// broadcaster 用于广播消息
func broadcaster() {
	// clients := make(map[Client]bool) // 所有连接的客户端

	// TODO: 实现广播逻辑
	// 1. 监听 messages, entering, leaving 通道
	// 2. 将消息广播给所有客户端
	// 3. 处理用户进入和离开
}

// handleConn 处理单个 TCP 连接
func handleConn(conn net.Conn) {
	// ch := make(chan string) // 对外发送消息的通道
	// go clientWriter(conn, ch)

	// TODO: 实现连接处理逻辑
	// 1. 获取客户端地址作为名字
	// 2. 将新用户发送到 entering 通道
	// 3. 使用 bufio.Scanner 读取用户输入并发送到 messages 通道
	// 4. 用户断开连接后，发送到 leaving 通道并关闭连接
}

func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg) // 注意：忽略网络错误
	}
}

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	go broadcaster()

	log.Println("Chat server started on localhost:8000")

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}
