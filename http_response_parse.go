package main

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/textproto"
	"strconv"
	"strings"
)

type HttpResponse struct {
	Status     string // e.g. "200 OK"
	StatusCode int    // e.g. 200
	Proto      string // e.g. "HTTP/1.0"
	ProtoMajor int    // e.g. 1
	ProtoMinor int    // e.g. 0
	Header     http.Header
	Body       []byte
}

func HttpResponseParse(res []byte) (*HttpResponse, error) {

	bodyIndex := bytes.Index(res, []byte("\r\n\r\n"))
	if bodyIndex == -1 {
		return nil, errors.New("body not found")
	}
	r := bytes.NewReader(res)
	lr := io.LimitReader(r, int64(len(res)))
	br := bufio.NewReader(lr)

	tp := textproto.NewReader(br)
	resp := new(HttpResponse)

	// Parse the first line of the response.
	line, err := tp.ReadLine()
	if err != nil {
		if err == io.EOF {
			err = io.ErrUnexpectedEOF
		}
		return nil, err
	}
	if i := strings.IndexByte(line, ' '); i == -1 {
		return nil, errors.New("malformed HTTP response")
	} else {
		resp.Proto = line[:i]
		resp.Status = strings.TrimLeft(line[i+1:], " ")
	}

	statusCode := resp.Status
	if i := strings.IndexByte(resp.Status, ' '); i != -1 {
		statusCode = resp.Status[:i]
	}
	if len(statusCode) != 3 {
		return nil, errors.New("malformed HTTP status code")
	}
	resp.StatusCode, err = strconv.Atoi(statusCode)
	if err != nil || resp.StatusCode < 0 {
		return nil, errors.New("malformed HTTP status code")
	}

	var ok bool
	if resp.ProtoMajor, resp.ProtoMinor, ok = http.ParseHTTPVersion(resp.Proto); !ok {
		return nil, errors.New("malformed HTTP version")
	}

	// Parse the response headers.
	mimeHeader, err := tp.ReadMIMEHeader()
	if err != nil {
		if err == io.EOF {
			err = io.ErrUnexpectedEOF
		}
		return nil, err
	}
	resp.Header = http.Header(mimeHeader)
	resp.Body = res[bodyIndex+4:]
	return resp, nil
}

func main() {
	res := "HTTP/1.1 200\r\nServer: SwooleServer\r\nContent-Type: text/html;charset=utf8\r\nContent-Length: 9\r\n\r\ntest-body"
	hp, err := HttpResponseParse([]byte(res))
	fmt.Println(hp, err)
	fmt.Println(string(hp.Body))
}
