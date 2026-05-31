package main

import (
	"flag"
	"fmt"
	"io"
	"net"
	"os"
	"time"
)

func main() {
	timeout := flag.Duration(
		"timeout",
		10*time.Second,
		"connection timeout",
	)

	flag.Parse()

	if flag.NArg() != 2 {
		fmt.Println("usage: telnet [--timeout=10s] host port")
		os.Exit(1)
	}

	host := flag.Arg(0)
	port := flag.Arg(1)

	address := net.JoinHostPort(host, port)

	conn, err := net.DialTimeout(
		"tcp",
		address,
		*timeout,
	)
	if err != nil {
		fmt.Println("connection error:", err)
		os.Exit(1)
	}

	defer conn.Close()

	fmt.Println("Connected to", address)

	done := make(chan struct{})

	go func() {
		io.Copy(os.Stdout, conn)
		close(done)
	}()

	go func() {
		io.Copy(conn, os.Stdin)
		conn.Close()
	}()

	<-done

	fmt.Println("\nConnection closed")
}