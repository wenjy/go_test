package main

import "fmt"

func main() {
	str1 := "aaa\n"
	fmt.Println(str1)
	str2 := `aaa\n`
	fmt.Println(str2)
	fmt.Println(str2[len(str2)-1]) // 打印最后一字节，这种转换方案只对纯 ASCII 码的字符串有效
	// &str[i] 非法
	str3 := str1 + str2
	fmt.Println(str3)
}
