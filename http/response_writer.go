package http

import (
	"fmt"
	"net"
)

const (
	Protocal = "HTTP/1.1"
	CRLF     = "\r\n"
)

const (
	StatusOK       = 200
	StatusNotFound = 404
)

const (
	OK       = "OK"
	NOTFOUND = "Not Found"
)

type ResponseWriter interface {
	Write([]byte) (int, error)
	WriteHeader(statusCode int)
	WriteHeaders(map[string]interface{})
}

type response struct {
	conn net.Conn
}

func NewResponse(conn net.Conn) *response {
	return &response{conn}
}

func (r *response) WriteHeader(statusCode int) {
	var statusString string
	switch statusCode {
	case StatusOK:
		statusString = OK
	case StatusNotFound:
		statusString = NOTFOUND
	}
	header := fmt.Sprintf("%s %d %s%s", Protocal, statusCode, statusString, CRLF)
	r.conn.Write([]byte(header))
}

func (r *response) WriteHeaders(headers map[string]interface{}) {
	for key, val := range headers {
		r.conn.Write([]byte(fmt.Sprintf("%s: %v%s", key, val, CRLF)))
	}
	r.conn.Write([]byte(CRLF))
}

func (r *response) Write(barr []byte) (int, error) {
	return r.conn.Write(barr)
}
