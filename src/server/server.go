package main

import (
	"fmt"
	"net"
)

func process(conn net.Conn) {
	defer conn.Close()

	for {
		buf := make([]byte, 1024)
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Printf("客户端退出 err=%v", err)
			return //!!!
		}
		fmt.Println(string(buf[:n]))
	}
}

func main() {
	fmt.Println("服务器开始监听...\n")
	// 监听本地的8888端口，适用于ipv4和ipv6
	listener, err := net.Listen("tcp", "0.0.0.0:8888")
	if err != nil {
		fmt.Println("listen err", err)
		return
	}
	defer listener.Close()

	for {
		// 等待客户端连接，可以使用telnet 0.0.0.0 8888 连接
		fmt.Println("等待客户端连接...\n")
		conn, err := listener.Accept()
		if err != nil {
			// 如果一次连接失败了，不必return，可以继续等待下一次连接
			fmt.Println("accept err", err)
		} else {
			fmt.Printf("accept suc, conn is %v\n", conn.RemoteAddr().String())
		}

		go process(conn)
	}
}
