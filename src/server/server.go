package server

import (
	"fmt"
	"net"
)

func main() {
	fmt.Println("服务器开始监听...")
	// 监听本地的8888端口，适用于ipv4和ipv6
	listener, err := net.Listen("tcp", "0.0.0.0:8888")


}
