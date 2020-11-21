package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	reqStr := "GET / HTTP/1.1\r\nHost: www.baidu.com\r\nContent-LengthProxy-Connection: keep-alive\r\nAccept: */*\r\nUser-Agent: Mozilla/5.0\r\n\r\ntest-body"

	index := strings.Index(reqStr, "\r\n\r\n")
	fmt.Println(index)
	fmt.Println(string([]byte(reqStr)[index+4:]))

	r := strings.NewReader(reqStr)
	lr := io.LimitReader(r, int64(len(reqStr)))
	br := bufio.NewReader(lr)
	request, err := http.ReadRequest(br)

	fmt.Println(request)

	fmt.Println(ioutil.ReadAll(request.Body))
	fmt.Println(err)

	res := "HTTP/1.1 200\r\nServer: SwooleServer\r\nContent-Type: text/html;charset=utf8\r\nContent-Length: 9\r\n\r\ntest-body"
	index = strings.Index(res, "\r\n\r\n")
	fmt.Println(index)
	fmt.Println(string([]byte(res)[index+4:]))

	r1 := bytes.NewReader([]byte(res))
	lr = io.LimitReader(r1, int64(len(res)))
	br = bufio.NewReader(lr)
	response, err := http.ReadResponse(br, request)

	fmt.Println(response)
	body, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(body))
	fmt.Println(err)
}
