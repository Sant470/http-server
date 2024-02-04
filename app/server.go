package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"

	http "github.com/codecrafters-io/http-server-starter-go/http"
)

func handleConn(conn net.Conn) {
	defer conn.Close()
	req, err := http.ReadRequest(bufio.NewReader(conn))
	if err != nil {
		log.Fatal("error reading request: ", err)
	}
	var rw http.ResponseWriter = http.NewResponse(conn)
	switch req.Path {
	case "/":
		res := fmt.Sprintf("%s %d %s%s%s", http.Protocal, http.StatusOK, "OK", http.CRLF, http.CRLF)
		if _, err := rw.Write([]byte(res)); err != nil {
			log.Fatal("error occurred while sending response: ", err)
		}
	default:
		res := fmt.Sprintf("%s %d %s%s%s", http.Protocal, http.StatusNotFound, "Not Found", http.CRLF, http.CRLF)
		if _, err := rw.Write([]byte(res)); err != nil {
			log.Fatal("error occurred while sending response: ", err)
		}
	}
}

func main() {
	l, err := net.Listen("tcp", "localhost:4221")
	if err != nil {
		fmt.Println("Failed to bind to port 4221")
		os.Exit(1)
	}
	conn, err := l.Accept()
	if err != nil {
		fmt.Println("Error accepting connection: ", err.Error())
		os.Exit(1)
	}
	handleConn(conn)
}
