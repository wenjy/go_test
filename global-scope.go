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
   a = "O" // 赋值给全局变量
   print(a)
}

// GOO