package step4_transport_layer

import (
	"context"
	"fmt"
	"net"
	"time"
)

// simulate tcp three-way handshake  模拟tcp三次握手

func tcpServer(ctx context.Context) {
	serverAddr := "127.0.0.1:8888"
	//创建服务器监听
	listener, err := net.Listen("tcp", serverAddr)
	if err != nil {
		return
	}
	conn, err := listener.Accept() //accept client
	defer conn.Close()
	if err != nil {
		return
	}
	conn.SetDeadline(time.Now().Add(100 * time.Second))
	_, err = conn.Write([]byte("SYN")) //first handshake
	if err != nil {
		return
	}
	buf := make([]byte, 1024)       //这里边如果是缓存区 就是会造成黏包情况呢  mysql服务器这里边也是会存在上述这种情况的吧
	readLine, err := conn.Read(buf) //receive client send message  second handshake
	if err != nil {
		return
	}
	message := buf[:readLine] //接收到客户端 SYN-ACK包机制

	fmt.Println(message)
	_, err = conn.Write([]byte("ACK"))
	if err != nil {
		return
	}
}

func tcpClient(ctx context.Context) {

}
