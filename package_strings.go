package main

import (
    "fmt"
    "strings"
)

func main() {
    var str string = "This is an example of a string"
    fmt.Printf("T/F? Does the string \"%s\" have prefix %s? ", str, "Th")
	fmt.Printf("%t\n", strings.HasPrefix(str, "Th"))
	fmt.Println()
	fmt.Printf("T/F? Does the string \"%s\" have suffix %s? ", str, "ing")
	fmt.Printf("%t\n", strings.HasSuffix(str, "ing"))
	fmt.Println()
	fmt.Printf("T/F? Does the string \"%s\" have Contains %s? ", str, "example")
	fmt.Printf("%t\n", strings.Contains(str, "example"))
	
	fmt.Println()
	fmt.Printf("T/F? Does the string \"%s\" Index %s? ", str, "example")
	fmt.Printf("%d\n", strings.Index(str, "example"))
	// -1 表示字符串不包含
	fmt.Println()
	fmt.Printf("T/F? Does the string \"%s\" LastIndex %s? ", str, "a")
	fmt.Printf("%d\n", strings.Index(str, "a"))

	str_zh := "我爱GO"
	fmt.Println()
	fmt.Printf("T/F? Does the string \"%s\" IndexRune %s? ", str, "爱")
	fmt.Printf("%d\n", strings.IndexRune(str_zh, '爱')) // 单引号表示rune 不做任何转义的原始内容

	// 统计
	fmt.Println()
	fmt.Printf("Number of G's in %s is: ", str_zh)
	fmt.Printf("%d\n", strings.Count(str_zh, "G"))
	
	// 替换
	replace_str := strings.Replace(str, "an", "AN", -1)
	fmt.Println(replace_str)

	// 重复生成
	repeat_str := strings.Repeat("a", 5)
	fmt.Println(repeat_str)

	// 大小写转换
	fmt.Println(strings.ToLower(str))
	fmt.Println(strings.ToUpper(str))

}
