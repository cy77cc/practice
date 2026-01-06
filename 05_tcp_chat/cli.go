package main

import (
	"fmt"
	"io"
	"os"
	"net"
	"sync"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	var wg sync.WaitGroup

	wg.Go(func() {
		defer wg.Done()
		if _, err := io.Copy(os.Stdout, conn); err != nil && err != io.EOF {
			fmt.Fprintf(os.Stderr, "read error: %v\n", err)
		}
	})

	wg.Go(func() {
		defer wg.Done()
		if _, err := io.Copy(conn, os.Stdin); err != nil && err != io.EOF {
			fmt.Fprintf(os.Stderr, "write error: %v\n", err)
		}
	})

	wg.Wait()
}
