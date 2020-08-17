package main

var a = "G"

func main() {
	n()
	m()
	n()
}

func n() {
	print(a)
}

func m() {
	a := "O" // 局部初始化变量并赋值
	print(a)
}

// GOG
