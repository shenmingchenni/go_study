package main

import (
	"fmt"
	"net"
)

type Server struct {
	IP   string
	Port int
}

func NewServer(ip string, port int) *Server {
	server := &Server{
		IP:   ip,
		Port: port,
	}
	return server
}
func (s *Server) Handle(conn net.Conn) {
	fmt.Println("链接建立成功")
	buf := make([]byte, 4096)
	for {
		n, err := conn.Read(buf)
		if n == 0 {
			fmt.Println("用户下线了")
			return
		}
		if err != nil {
			fmt.Println("conn.Read err:", err)
		}
		message := string(buf[:n-1])
		fmt.Printf("收到消息:%s\n", message)
	}
}
func (s *Server) Start() {
	Listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", s.IP, s.Port))
	if err != nil {
		fmt.Println("net.Listen err:", err)
		return
	}
	defer Listener.Close()
	for {
		Conn, err := Listener.Accept()
		if err != nil {
			fmt.Println("Listener.accept err:", err)
			continue
		}
		go s.Handle(Conn)
	}
}
