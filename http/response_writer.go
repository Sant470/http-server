package http

import (
	"net"
)

const (
	Protocal = "HTTP/1/1"
	CRLF     = "\r\n"
)

const (
	StatusOK = 200
)

type ResponseWriter interface {
	Write([]byte) (int, error)
}

type response struct {
	conn net.Conn
}

func NewResponse(conn net.Conn) *response {
	return &response{conn}
}

func (r *response) Write(barr []byte) (int, error) {
	return r.conn.Write(barr)
}
