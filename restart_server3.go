// server.go
package main

import (
	"flag"
	"log"
	"net"
	"os"
	"os/exec"
	"os/signal"
	"path/filepath"
	"sync"
	"syscall"
	"time"
)

var (
	workSpace      string
	logger         *log.Logger
	writeTimeout   = time.Second * 5
	readTimeout    = time.Second * 5
	signalChan     = make(chan os.Signal)
	connFiles      sync.Map
	serverListener net.Listener
	isUpdate       = false
)

func init() {
	flag.StringVar(&workSpace, "w", ".", "Usage:\n ./server -w=workspace")
	flag.Parse()
	file, err := os.OpenFile(filepath.Join(workSpace, "server.log"), os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0777)
	if err != nil {
		panic(err)
	}
	logger = log.New(file, "", 11)
	go signalHandler()
}

func main() {
	var err error
	serverListener, err = net.Listen("tcp", ":7000")
	if err != nil {
		panic(err)
	}
	for {
		// 更新状态不接受新连接
		if isUpdate == true {
			continue
		}
		conn, err := serverListener.Accept()
		if err != nil {
			logger.Println("conn error")
			continue
		}
		c := conn.(*net.TCPConn)
		go connectionHandler(c)
	}
}

// 处理连接
func connectionHandler(conn *net.TCPConn) {
	file, _ := conn.File()
	connFiles.Store(file, true)
	logger.Printf("conn fd %d\n", file.Fd())

	defer func() {
		connFiles.Delete(file)
		_ = conn.Close()
	}()

	for {
		// 更新状态不处理连接
		if isUpdate == true {
			continue
		}
		err := conn.SetReadDeadline(time.Now().Add(readTimeout))
		if err != nil {
			logger.Println(err.Error())
			return
		}
		rBuf := make([]byte, 4)
		_, err = conn.Read(rBuf)
		if err != nil {
			logger.Println(err.Error())
			return
		}
		if string(rBuf) != "ping" {
			logger.Println("failed to parse the message " + string(rBuf))
			return
		}
		err = conn.SetWriteDeadline(time.Now().Add(writeTimeout))
		if err != nil {
			logger.Println(err.Error())
			return
		}
		_, err = conn.Write([]byte(`pong`))
		if err != nil {
			logger.Println(err.Error())
			return
		}
	}
}

// 信号
func signalHandler() {
	signal.Notify(signalChan, syscall.SIGUSR2)
	for {
		sc := <-signalChan
		switch sc {
		case syscall.SIGUSR2:
			gracefulExit()
		default:
			continue
		}
	}
}

// 重启
func gracefulExit() {
	// 先更新状态
	isUpdate = true

	var err error

	var files = make([]*os.File, 1)
	lfile, _ := serverListener.(*net.TCPListener).File()
	files[0] = lfile

	// 所有连接发送到 unix sock
	connFiles.Range(func(key, value interface{}) bool {
		if key == nil || value == nil {
			return false
		}
		file := key.(*os.File)
		defer func() {
			_ = file.Close()
		}()
		files = append(files, file)
		return true
	})

	// log.Println(files)
	path := os.Args[0]
	var args []string
	if len(os.Args) > 1 {
		args = os.Args[1:]
	}

	cmd := exec.Command(path, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.ExtraFiles = files
	cmd.Env = os.Environ()

	err = cmd.Start()
	if err != nil {
		log.Fatalf("Restart: Failed to launch, error: %v", err)
	}
	os.Exit(0)
}
