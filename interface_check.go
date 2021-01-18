package main

type Checker interface {
	Checker1()
	Checker2()
}

type Check struct {
}

// 检测是否实现接口
var _ Checker = (*Check)(nil)

func main() {

}

func (c *Check) Checker1() {

}

func (c *Check) Checker2() {

}
