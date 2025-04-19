package main

import (
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	ln, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		go handleConnection(conn)
	}
}
func handleConnection(conn net.Conn) {
	defer conn.Close()
	buf := make([]byte, 1024)

	_, err := conn.Read(buf)
	if err != nil {
		fmt.Println(err)
		return
	}

	bufStr := strings.Fields(string(buf))
	if bufStr[1] == "/" || bufStr[1] == "/index.html" {
		err = serveRoot(bufStr, conn)
		if err != nil {
			fmt.Println(err)
		}
	} else {
		responseStr := bufStr[2] + " 404 Not Found\r\n\r\n"
		conn.Write([]byte(responseStr))
	}
}

func serveRoot(bufStr []string, conn net.Conn) error {
	htmlFile, err := os.ReadFile("www/index.html")
	if err != nil {
		return err
	}
	responseStr := bufStr[2] + " 200 OK\r\n\r\n"
	conn.Write(append([]byte(responseStr), htmlFile...))
	return nil
}
