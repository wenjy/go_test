package main

import (
	"crypto/tls"
	"fmt"
	"net"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/valyala/fasthttp"
)

var AppPath string

type tcpKeepaliveListener struct {
	*net.TCPListener
	keepalive       bool
	keepalivePeriod time.Duration
}

func (ln tcpKeepaliveListener) Accept() (net.Conn, error) {
	tc, err := ln.AcceptTCP()
	if err != nil {
		return nil, err
	}
	if err := tc.SetKeepAlive(ln.keepalive); err != nil {
		tc.Close() //nolint:errcheck
		return nil, err
	}
	if ln.keepalivePeriod > 0 {
		if err := tc.SetKeepAlivePeriod(ln.keepalivePeriod); err != nil {
			tc.Close() //nolint:errcheck
			return nil, err
		}
	}
	return tc, nil
}

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

	tlsConfig := &tls.Config{
		PreferServerCipherSuites: true,
	}

	cert, err := tls.LoadX509KeyPair(certFile, keyFile)
	if err != nil {
		fmt.Println(fmt.Errorf("cannot load TLS key pair from certFile=%q and keyFile=%q: %s", certFile, keyFile, err))
	}

	tlsConfig.Certificates = append(tlsConfig.Certificates, cert)

	for port := 10000; port <= 11000; port++ {
		go func(port int) {
			/* err := server.ListenAndServeTLS(":"+strconv.Itoa(port), certFile, keyFile)
			if err != nil {
				fmt.Println(err)
			} */
			ln, err := net.Listen("tcp4", ":"+strconv.Itoa(port))
			if err != nil {
				fmt.Println(err)
				return
			}
			if tcpln, ok := ln.(*net.TCPListener); ok {
				err = server.Serve(
					tls.NewListener(tcpKeepaliveListener{
						TCPListener:     tcpln,
						keepalive:       server.TCPKeepalive,
						keepalivePeriod: server.TCPKeepalivePeriod,
					}, tlsConfig),
				)
				if err != nil {
					fmt.Println(err)
					return
				}
			}
		}(port)
	}

	for {

	}

}
