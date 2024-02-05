package http

import (
	"fmt"
	"io"
	"net"
)

const (
	Protocal = "HTTP/1.1"
	CRLF     = "\r\n"
	//status code
	StatusOK            = 200
	StatusCreated       = 201
	StatusNotFound      = 404
	InternalServerError = 500
	// string status
	OK                  = "OK"
	NOTFOUND            = "Not Found"
	CREATED             = "Created"
	INTERNALSERVERERROR = "Internal Server Error"
)

type ResponseWriter interface {
	Write(io.Reader) (int64, error)
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
	case StatusCreated:
		statusString = CREATED
	case InternalServerError:
		statusString = INTERNALSERVERERROR
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

func (r *response) Write(reader io.Reader) (int64, error) {
	return io.Copy(r.conn, reader)
}
