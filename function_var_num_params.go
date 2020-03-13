package main

import "fmt"

func main() {
    x := min(1, 3, 2, 0)
    fmt.Println("The minimum is: ", x)
    //切片
    slice := []int{7,9,3,5,1}
    x = min(slice...)
    fmt.Println("The minimum in the slice is: ", x)

    a := 1
    b := "b"
    c := 1.1
    d := false
    typecheck(a,b,c,d)
}

func min(s ...int) int {
    if len(s)==0 {
        return 0
    }
    min := s[0]
    for _, v := range s {
        if v < min {
            min = v
        }
    }
    return min
}

func typecheck(values ...interface{}) {
    for _, value := range values {
        switch value.(type) {
            case int: fmt.Println(value, "int")
            case float64: fmt.Println(value, "float64")
            case string: fmt.Println(value, "string")
            case bool: fmt.Println(value, "bool")
            default: fmt.Println(value, "unkown")
        }
    }
}
