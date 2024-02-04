package http

import (
	"bufio"
	"strings"
)

type Request struct {
	Method      string
	Path        string
	HTTPVersion string
}

func ReadRequest(r *bufio.Reader) (*Request, error) {
	req := new(Request)
	barr, _ := r.ReadBytes('\n')
	startLine := strings.Split(string(barr), " ")
	req.Method = startLine[0]
	req.Path = startLine[1]
	req.HTTPVersion = startLine[2]
	return req, nil
}
