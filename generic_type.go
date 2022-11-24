package main

import "fmt"

// 类型约束限定了只能使用 int 或 float32 或 float64 来实例化自己
type Slice[T int | float32 | float64] []T

// MyMap类型定义了两个类型形参 KEY 和 VALUE。分别为两个形参指定了不同的类型约束
// 这个泛型类型的名字叫： MyMap[KEY, VALUE]
type MyMap[KEY int | string, VALUE float32 | float64] map[KEY]VALUE

// 一个泛型类型的结构体。可用 int 或 sring 类型实例化
type MyStruct1[T int | string] struct {
	Name string
	Data T
}

// 一个泛型接口
type IPrintData[T int | float32 | string] interface {
	Print(data T)
}

// 一个泛型通道，可用类型实参 int 或 string 实例化
type MyChan1[T int | string] chan T

// 类型形参是可以互相套用
type WowStruct[T int | float32, S []T] struct {
	Data     S
	MaxValue T
	MinValue T
}

// 指定底层类型
type Int1 interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

type Uint1 interface {
	~uint | ~uint8 | ~uint16 | ~uint32
}
type Float1 interface {
	~float32 | ~float64
}

type Slice1[T Int1 | Uint1 | Float1] []T

type MyInt int

// 空接口代表所有类型的集合。写入类型约束意味着所有类型都可拿来做类型实参
type Slice2[T interface{}] []T

// Go1.18开始提供了一个和空接口 interface{} 等价的新关键词 any
type Slice3[T any] []T // 代码等价于 type Slice[T interface{}] []T

// 泛型接口
type DataProcessor[T any] interface {
	Process(oriData T) (newData T)
	Save(data T) error
}

// 约束了实现接口的类型
type DataProcessor2[T any] interface {
	int | ~struct{ Data interface{} }

	Process(data T) (newData T)
	Save(data T) error
}

func main() {
	// 这里传入了类型实参int，泛型类型Slice[T]被实例化为具体的类型 Slice[int]
	var a Slice[int] = []int{1, 2, 3}
	fmt.Printf("Type Name: %T\n", a) //输出：Type Name: Slice[int]

	fmt.Println(a.Sum())

	// 传入类型实参float32, 将泛型类型Slice[T]实例化为具体的类型 Slice[string]
	var b Slice[float32] = []float32{1.0, 2.0, 3.0}
	fmt.Printf("Type Name: %T\n", b) //输出：Type Name: Slice[float32]

	// 用类型实参 string 和 flaot64 替换了类型形参 KEY 、 VALUE，泛型类型被实例化为具体的类型：MyMap[string, float64]
	var map111 MyMap[string, float64] = map[string]float64{
		"jack_score": 9.6,
		"bob_score":  8.4,
	}

	fmt.Printf("Type Name: %T\n", map111) // MyMap[string,float64]

	// 泛型类型 WowStuct[T, S] 被实例化后的类型名称就叫 WowStruct[int, []int]
	var ws11 WowStruct[int, []int] = WowStruct[int, []int]{
		Data:     []int{1, 2, 3},
		MaxValue: 3,
		MinValue: 1,
	}
	fmt.Printf("Type Name: %T\n", ws11) // WowStruct[int,[]int]

	a11 := add11[int](1, 2) // 传入类型实参int，计算结果为 3
	fmt.Println(a11)
	a11 = add11(1, 2)
	fmt.Println(a11)
	a22 := add11[float32](1.0, 2.0) // 传入类型实参float32, 计算结果为 3.0
	fmt.Println(a22)
	//a22 = add11(1.0, 2.0) // 这样有问题，默认float64

	var s Slice[int] // 正确
	fmt.Println(s)

	var s2 Slice1[MyInt] // MyInt底层类型是int，所以可以用于实例化
	fmt.Println(s2)
}

func (s Slice[T]) Sum() T {
	var sum T
	for _, value := range s {
		sum += value
	}
	return sum
}

// 乏型函数
func add11[T int | float32 | float64](a T, b T) T {
	return a + b
}
