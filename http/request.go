package http

import (
	"bufio"
	"strings"
)

type Request struct {
	Method      string
	URL         string
	HTTPVersion string
	Headers     map[string]string
	reader      *bufio.Reader
}

func (r *Request) readHeaders() {
	for {
		barr, _ := r.reader.ReadBytes('\n')
		els := strings.Split(string(barr), ":")
		if len(els) < 2 {
			break
		}
		r.Headers[els[0]] = strings.Join(els[1:], ":")
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
	return req, nil
}
