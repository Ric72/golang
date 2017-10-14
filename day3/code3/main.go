// code3 project main.go
//批量创建map集合
package main

import (
	"fmt"
)

type Product struct {
	X, Y, Z int
}

func main() {
	var m = map[string]Product{
		"a": {1, 2, 3},
		"b": {4, 5, 6},
		"c": {7, 8, 9},
	}

	m["a"] = Product{18, 3, 3}

	fmt.Println(m["a"].X)
}
