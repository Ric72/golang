// code15 project main.go
//对空的slice赋值
package main

import (
	"fmt"
)

func main() {
	slice := make([]int, 5)
	for v := range slice {
		slice[v] = v + 1
	}
	fmt.Println(slice)
}
