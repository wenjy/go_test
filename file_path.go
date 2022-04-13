package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	fmt.Println(filepath.FromSlash("a/n/c"))
	fmt.Println(filepath.FromSlash("a\\n\\c"))
}
