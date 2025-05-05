package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	fmt.Println("Logs from your program will appear here!")
	s := Server{}
	s.Start()

}

type Server struct {
	listener net.Listener
}

func (s *Server) Start() {
	s.Listen()
	defer s.Close()

	for {
		conn := s.Accept()
		fmt.Println("Accepted connection from: ", conn.RemoteAddr())
		go s.handleConnection(conn)
	}

}
func (s *Server) Listen() {

	l, err := net.Listen("tcp", "0.0.0.0:4221")
	if err != nil {
		fmt.Println("Failed to bind to port 4221")
		os.Exit(1)
	}
	s.listener = l
}

func (s *Server) Accept() net.Conn {
	conn, err := s.listener.Accept()
	if err != nil {
		fmt.Println("Error accepting connection: ", err.Error())
		os.Exit(1)
	}
	return conn
}
func (s *Server) Close() {
	err := s.listener.Close()
	if err != nil {
		fmt.Println("Failed to close listener:", err.Error())
	}
}
func (s *Server) handleConnection(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)
	line, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading request:", err)
		conn.Close()
		return
	}
	parts := strings.Split(line, " ")
	path := parts[1]
	if len(parts) < 2 {
		conn.Write([]byte("HTTP/1.1 400 Bad Request\r\n\r\n"))
		conn.Close()
		return
	}
	switch path {
	case "/":
		fmt.Println(parts)
		conn.Write([]byte("HTTP/1.1 200 OK\r\n\r\nWelcome!\n"))
	case "/hello":
		conn.Write([]byte("HTTP/1.1 200 OK\r\n\r\nHello, World!\n"))
	default:
		conn.Write([]byte("HTTP/1.1 404 Not Found\r\n\r\n"))
	}
}
