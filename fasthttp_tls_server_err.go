package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"

	"github.com/valyala/fasthttp"
)

var AppPath string

func main() {
	var err error
	if AppPath, err = filepath.Abs(filepath.Dir(os.Args[0])); err != nil {
		panic(err)
	}

	server := &fasthttp.Server{
		Handler: func(ctx *fasthttp.RequestCtx) {

		},
	}

	certFile := AppPath + "/ssl/test.crt"
	keyFile := AppPath + "/ssl/test.key"

	for port := 10000; port <= 10020; port++ {
		go func(port int) {
			err := server.ListenAndServeTLS(":"+strconv.Itoa(port), certFile, keyFile)
			if err != nil {
				fmt.Println(err)
			}
		}(port)
	}

	for {

	}

}
