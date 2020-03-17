package main

import (
    "./person"
    "fmt"
)

func main() {
	p := new(person.Person)
	// fmt.Println(p.firstName)
    // p.firstName undefined (cannot refer to unexported field or method firstName)
    // p.firstName = "Eric"
    p.SetFirstName("Eric")
    fmt.Println(p.FirstName()) // Output: Eric
}