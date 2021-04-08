package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8888")
	if err != nil {
		fmt.Println("client dial err=", err)
		return
	}

	// 客户端可以从终端发送多行数据给服务器
	reader := bufio.NewReader(os.Stdin)
	for {
		// 从终端读取一行用户输入，并准备发送给服务器
		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("read string err =", err)
		}
		// 如果用户输入exit就退出
		line = strings.Trim(line, "\n")
		if line == "exit" {
			break
		}

		_, err = conn.Write([]byte(line))
		if err != nil {
			fmt.Println("conn.Write err =", err)
		}

	}
}
