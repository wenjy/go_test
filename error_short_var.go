package main

import "fmt"

func main() {
	err := shadow()
	fmt.Println(err)
}

func shadow() (err error) {
	x, err := check1() // x是新创建变量，err是被赋值
	if err != nil {
		return // 正确返回err
	}
	if y, err := check2(x); err != nil { // y和if语句中err被创建
		return // if语句中的err覆盖外面的err，所以错误的返回nil！ err is shadowed during return
	} else {
		fmt.Println(y)
	}
	return
}

func check1() (int, error) {
	return 1, fmt.Errorf("%s", "error1")
}

func check2(i int) (int, error) {
	return i, fmt.Errorf("%s", "error2")
}
