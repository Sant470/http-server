package http

import (
	"bufio"
)

type Request struct{}

func ReadRequest(r *bufio.Reader) (*Request, error) {
	req := new(Request)
	r.ReadBytes('\n')
	return req, nil
}
