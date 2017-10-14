// code2 project main.go
//map的值为一个结构体
package main

import (
	"fmt"
)

func main() {
	type Product struct {
		name   string
		price  int
		amount int
	}
	m := make(map[string]Product)
	m["basketball"] = Product{"篮球", 123, 50}
	delete(m, "basketball")
	fmt.Println(m)
}
