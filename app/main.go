package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

// Ensures gofmt doesn't remove the "net" and "os" imports above (feel free to remove this!)
var _ = net.Listen
var _ = os.Exit

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
	if len(parts) < 2 {
		conn.Write([]byte("HTTP/1.1 400 Bad Request\r\n\r\n"))
		conn.Close()
		return
	}
	if parts[1] == "/" {
		conn.Write([]byte("HTTP/1.1 200 OK\r\n\r\n"))
	} else {

		conn.Write([]byte("HTTP/1.1 404 Not Found\r\n\r\n"))
	}

}
