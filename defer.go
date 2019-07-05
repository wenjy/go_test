package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// defer 语句会延迟函数的执行直到上层函数返回。
// 延迟调用的参数会立刻生成，但是在上层函数返回前函数都不会被调用。
// 延迟的函数调用被压入一个栈中。当函数返回时， 会按照后进先出的顺序调用被延迟的函数调用。
func main() {
	defer fmt.Println("world")

	fmt.Println("hello")

	fmt.Println("counting")

	for i := 0; i < 3; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")

	// 写入文件
	currentDir := GetCurrentDirectory()
	file,err1 := os.Create(currentDir + "/file")
	if err1 != nil {
		fmt.Println(err1)
	}
	fileStr1 := []byte("123456我是谁")

	n, err2 := file.Write(fileStr1)
	if err2 != nil {
		fmt.Println(err1)
	}
	fmt.Println(n)
	err3 := file.Close()
	if err3 != nil {
		fmt.Println(err1)
	}

	fileStr, err := Contents(currentDir + "/file")

	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Printf("fileStr:%s\n", fileStr)
}


// Contents returns the file's contents as a string.
func Contents(filename string) (string, error) {
	f, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer f.Close() // f.Close will run when we're finished.

	var result []byte
	buf := make([]byte, 100)
	for {
		n, err := f.Read(buf[0:])
		result = append(result, buf[0:n]...) // append is discussed later.
		if err != nil {
			if err == io.EOF {
				break
			}
			return "", err // f will be closed if we return here.
		}
	}
	return string(result), nil // f will be closed if we return here.
}

func GetCurrentDirectory() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0])) //返回绝对路径  filepath.Dir(os.Args[0])去除最后一个元素的路径
	if err != nil {
		log.Fatal(err)
	}
	return strings.Replace(dir, "\\", "/", -1) //将\替换成/
}
