package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"

	http "github.com/codecrafters-io/http-server-starter-go/http"
)

func handleConn(conn net.Conn) {
	defer conn.Close()
	req, err := http.ReadRequest(bufio.NewReader(conn))
	if err != nil {
		log.Fatal("error reading request: ", err)
	}
	var rw http.ResponseWriter = http.NewResponse(conn)
	_, query, _ := strings.Cut(req.URL, "/")
	path, query, _ := strings.Cut(query, "/")
	path = "/" + path
	fmt.Println("path: ", path)
	fmt.Println("query: ", query)
	switch path {
	case "/":
		rw.WriteHeader(http.StatusOK)
		rw.WriteHeaders(map[string]interface{}{})
	case "/echo":
		rw.WriteHeader(http.StatusOK)
		rw.WriteHeaders(map[string]interface{}{
			"Content-Type":   "text/plain",
			"Content-Length": len(query),
		})
		if _, err := rw.Write([]byte(query)); err != nil {
			log.Println("error writing body: ", err)
		}
	default:
		rw.WriteHeader(http.StatusNotFound)
		rw.WriteHeaders(map[string]interface{}{})
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
