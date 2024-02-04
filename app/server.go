package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"

	http "github.com/codecrafters-io/http-server-starter-go/http"
)

// HTTP/1.1 200 OK\r\n\r\n
func handleConn(conn net.Conn) {
	http.ReadRequest(bufio.NewReader(conn))
	var rw http.ResponseWriter = http.NewResponse(conn)
	res := fmt.Sprintf("%s %d %s%s%s", http.Protocal, http.StatusOK, "OK", http.CRLF, http.CRLF)
	if _, err := rw.Write([]byte(res)); err != nil {
		log.Fatal("error occurred while sending response: ", err)
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
