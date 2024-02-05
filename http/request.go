package http

import (
	"bufio"
	"strconv"
	"strings"
)

type Request struct {
	reader      *bufio.Reader
	Method      string
	URL         string
	HTTPVersion string
	Headers     map[string]string
	Body        []byte
}

func (r *Request) readBody() {
	length, _ := strconv.Atoi(r.Headers["Content-Length"])
	barr := make([]byte, length)
	r.reader.Read(barr)
	r.Body = barr
}

func (r *Request) readHeaders() {
	for {
		barr, _ := r.reader.ReadString('\n')
		els := strings.Split(string(barr), ":")
		if len(els) < 2 {
			break
		}
		r.Headers[els[0]] = strings.TrimSpace(strings.Join(els[1:], ":"))
	}
}

func ReadRequest(r *bufio.Reader) (*Request, error) {
	req := &Request{reader: r, Headers: make(map[string]string)}
	barr, _ := req.reader.ReadBytes('\n')
	startLine := strings.Split(string(barr), " ")
	req.Method = startLine[0]
	req.URL = startLine[1]
	req.HTTPVersion = startLine[2]
	req.readHeaders()
	if req.Method == "POST" {
		req.readBody()
	}
	return req, nil
}
