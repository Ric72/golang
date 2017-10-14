// code1 project main.go
package main

import (
	"fmt"
)

func main() {
	//创建一个map集合
	m := make(map[int]int)
	m[11] = 11
	m[12] = 12

	delete(m, 11)

	fmt.Println(m)
}
