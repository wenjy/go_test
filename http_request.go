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

	reqStrConnect := "CONNECT www.baidu.com:443 HTTP/1.1\r\nHost: www.baidu.com:443\r\nProxy-Connection: keep-alive\r\nUser-Agent: Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/76.0.3809.100 Safari/537.36\r\nContent-Length: 9\r\n\r\ntest-body"
	r2 := strings.NewReader(reqStrConnect)
	lr2 := io.LimitReader(r2, int64(len(reqStrConnect)))
	br2 := bufio.NewReader(lr2)
	request2, err := http.ReadRequest(br2)

	fmt.Println(request2)
	fmt.Println(ioutil.ReadAll(request2.Body))

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
