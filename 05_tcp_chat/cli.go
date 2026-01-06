package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"sync"
)

func main() {
	conn, err := net.Dial("tcp", ":8000")
	if err != nil {
		panic(err)
	}

	reader := bufio.NewReader(conn)
	writer := bufio.NewWriter(conn)

	var wg sync.WaitGroup

	wg.Go(func() {
		buf := make([]byte, 1024)
		n, err := reader.Read(buf)
		if err != nil {
			if err == io.EOF {
				return
			}
		}
		fmt.Println(string(buf[:n]))
	})

	wg.Go(func() {
		s := ""
		fmt.Scanf("%s", &s)
		writer.WriteString(s)
	})

	wg.Wait()
}
