package main

import (
	"fmt"
	"strings"
)

func main() {
	v1 := "v1.0-0"
	v2 := "v2.1-1-3"
	replaceMap := map[string]string{"V": "", "v": "", "-": "."}
	//keywords := {"alpha,beta,rc,p"}
	for k, v := range replaceMap {
		if strings.Contains(v1, k) {
			v1 = strings.Replace(v1, k, v, -1)
		}
		if strings.Contains(v2, k) {
			v2 = strings.Replace(v2, k, v, -1)
		}
	}

	fmt.Println(v1, v2)
}
