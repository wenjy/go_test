package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"strconv"
	"strings"

	"github.com/valyala/fasthttp"
)

func main() {
	reqStr := "GET / HTTP/1.1\r\nHost: www.baidu.com\r\nProxy-Connection: keep-alive\r\nAccept: */*\r\nUser-Agent: Mozilla/5.0\r\nContent-Length: 9\r\n\r\ntest-body"

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

	reqStrConnect := "CONNECT www.baidu.com:443 HTTP/1.1\r\nHost: www.baidu.com:443\r\nProxy-Connection: keep-alive\r\nConnection: keep-alive\r\nUser-Agent: Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/76.0.3809.100 Safari/537.36\r\n\r\n"
	reqStrConnect = "CONNECT 2021.ip138.com:443 HTTP/1.1\r\nHost: 2021.ip138.com:443\r\nProxy-Connection: keep-alive\r\nConnection: keep-alive\r\nUser-Agent: Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/76.0.3809.100 Safari/537.36\r\n\r\n"

	reqStrConnect = "CONNECT www.example.com:443 HTTP/1.1\r\n\r\n"
	r2 := strings.NewReader(reqStrConnect)
	lr2 := io.LimitReader(r2, int64(len(reqStrConnect)))
	br2 := bufio.NewReader(lr2)
	request2, err := http.ReadRequest(br2)

	host, port := parseHostPort(request2)
	fmt.Println(host, port)

	r11 := strings.NewReader(reqStrConnect)
	lr11 := io.LimitReader(r11, int64(len(reqStrConnect)))
	br11 := bufio.NewReader(lr11)
	request11 := fasthttp.AcquireRequest()
	request11.Read(br11)
	host, port = parseHostPort2(request11)
	fmt.Println(host, port)
	return
	request2.Header.Set("Connection", "Close")
	fmt.Println(request2)
	fmt.Println(request2.Host)
	fmt.Println(ioutil.ReadAll(request2.Body))

	// 测试 fasthttp
	// reqStrConnect = "CONNECT 2021.ip138.com:443 HTTP/1.1\r\nHost: server.example.com:80\r\n\r\n"
	reqStrConnect = "CONNECT 2021.ip138.com:443 HTTP/1.1\r\n\r\n"
	r10 := strings.NewReader(reqStrConnect)
	lr10 := io.LimitReader(r10, int64(len(reqStrConnect)))
	br10 := bufio.NewReader(lr10)
	request10 := fasthttp.AcquireRequest()
	request10.Read(br10)
	fmt.Println("request10 Host", string(request10.Header.Host()))
	fmt.Println("request10 Method", string(request10.Header.Method()))
	fmt.Println("request10 URI Host", string(request10.URI().Host()))
	fmt.Println("request10 URI Path", string(request10.URI().Path()))
	if request10.Header.IsConnect() {
		fmt.Println("request10 IsConnect")
	}
	hostStr, portStr, _ := net.SplitHostPort(strings.Trim(string(request10.URI().Path()), "/"))
	fmt.Println("request10 SplitHostPort", hostStr, portStr)

	res := "HTTP/1.1 200\r\nServer: SwooleServer\r\nContent-Type: text/html;charset=utf8\r\nContent-Length: 10\r\n\r\ntest-body"
	index = strings.Index(res, "\r\n\r\n")
	fmt.Println(index)
	fmt.Println(string([]byte(res)[index+4:]))

	r1 := bytes.NewReader([]byte(res))
	lr = io.LimitReader(r1, int64(len(res)))
	br = bufio.NewReader(lr)
	response, err := http.ReadResponse(br, &http.Request{})

	fmt.Println("response", response)
	body, err := ioutil.ReadAll(response.Body)
	fmt.Println(string(body))
	fmt.Println(err)
}

func parseHostPort(r *http.Request) (string, uint16) {
	defaultPort := uint16(80)
	host := r.Host
	hostStr, portStr, err := net.SplitHostPort(host)
	if err != nil {
		return host, defaultPort
	}
	port, err := strconv.Atoi(portStr)
	if err != nil {
		return hostStr, defaultPort
	}
	return hostStr, uint16(port)
}

func parseHostPort2(r *fasthttp.Request) (string, uint16) {
	defaultPort := uint16(80)
	host := string(r.Host())
	if host == "" && r.Header.IsConnect() {
		host = strings.Trim(string(r.URI().Path()), "/")
	}
	hostStr, portStr, err := net.SplitHostPort(host)
	if err != nil {
		return host, defaultPort
	}
	port, err := strconv.Atoi(portStr)
	if err != nil {
		return hostStr, defaultPort
	}
	return hostStr, uint16(port)
}
