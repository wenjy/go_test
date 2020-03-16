package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

// map 在使用之前必须用 make 而不是 new 来创建；值为 nil 的 map 是空的，并且不能赋值。

var m map[string]Vertex

// map 的文法跟结构体文法相似，不过必须有键名。
var m1 = map[string]Vertex{
	"Bell Labs": Vertex{
		40.68433, -74.39967,
	},
	"Google": Vertex{
		37.42202, -122.08408,
	},
}

// 如果顶级的类型只有类型名的话，可以在文法的元素中省略键名。
var m2 = map[string]Vertex{
	"Bell Labs": {40.68433, -74.39967},
	"Google":    {37.42202, -122.08408},
}


func main() {
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])
	fmt.Println(m1)
	fmt.Println(m2)

	m3 := make(map[string]int)

	// 在 map m 中插入或修改一个元素：
	m3["Answer"] = 42
	fmt.Println("The value:", m3["Answer"])

	m3["Answer"] = 48
	fmt.Println("The value:", m3["Answer"])

	// 删除元素
	delete(m3, "Answer")
	fmt.Println("The value:", m3["Answer"])

	// 通过双赋值检测某个键存在：
	// 如果 key 在 m 中，`ok` 为 true 。否则， ok 为 `false`，并且 elem 是 map 的元素类型的零值。
	v, ok := m3["Answer"]
	fmt.Println("The value:", v, "Present?", ok)
	m3["abc"] = 1;
	if _, ok := m3["abc"]; ok {
		fmt.Println("key abc in m3")
	}
}

