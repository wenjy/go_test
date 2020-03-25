package main
import (
    "bufio"
    "fmt"
    "os"
)

func main() {
    inputFile, inputError := os.Open("test_file.dat")
    if inputError != nil {
        fmt.Printf("An error occurred on opening the inputfile\n" +
            "Does the file exist?\n" +
            "Have you got acces to it?\n")
        return // exit the function on error
    }
    defer inputFile.Close()

	inputReader := bufio.NewReader(inputFile)
	// 带缓冲的读取
	buf := make([]byte, 1024)
	n, err := inputReader.Read(buf)
	if err == nil {
		fmt.Println(n)
	}
}