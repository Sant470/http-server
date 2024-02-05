package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"path/filepath"
	"strings"

	http "github.com/codecrafters-io/http-server-starter-go/http"
)

var dir string

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
		if _, err := rw.Write(strings.NewReader(query)); err != nil {
			log.Println("error writing body: ", err)
		}
	case "/user-agent":
		agent := req.Headers["User-Agent"]
		rw.WriteHeader(http.StatusOK)
		rw.WriteHeaders(map[string]interface{}{
			"Content-Type":   "text/plain",
			"Content-Length": len(agent),
		})
		if _, err := rw.Write(strings.NewReader(agent)); err != nil {
			log.Println("error writing body: ", err)
		}
	case "/files":
		if req.Method == "POST" {
			path := filepath.Join(dir, query)
			file, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
			if err != nil {
				rw.WriteHeader(http.InternalServerError)
				break
			}
			defer file.Close()
			if _, err := io.Copy(file, bytes.NewReader(req.Body)); err != nil {
				fmt.Println("error: ", err)
				rw.WriteHeader(http.InternalServerError)
				break
			}
			// _, err = file.Write(req.Body)
			// if err != nil {
			// 	fmt.Println("error: ", err)
			// }
			rw.WriteHeader(http.StatusCreated)
		}
		if req.Method == "GET" {
			path := filepath.Join(dir, query)
			fi, err := os.Stat(path)
			if err != nil {
				rw.WriteHeader(http.StatusNotFound)
				rw.WriteHeaders(map[string]interface{}{})
				break
			}
			file, _ := os.Open(path)
			defer file.Close()
			rw.WriteHeader(http.StatusOK)
			rw.WriteHeaders(map[string]interface{}{
				"Content-Type":   "application/octet-stream",
				"Content-Length": fi.Size(),
			})
			if _, err := rw.Write(file); err != nil {
				log.Println("error writing body: ", err)
			}
		}
	default:
		rw.WriteHeader(http.StatusNotFound)
		rw.WriteHeaders(map[string]interface{}{})
	}
}

func main() {
	flag.StringVar(&dir, "directory", "", "dir name")
	flag.Parse()
	l, err := net.Listen("tcp", "localhost:4221")
	if err != nil {
		fmt.Println("Failed to bind to port 4221")
		os.Exit(1)
	}
	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting connection: ", err.Error())
			os.Exit(1)
		}
		go handleConn(conn)
	}
}
