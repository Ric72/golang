// code1 project main.go
package main

import (
	"fmt"
)

//关于if语句的初始化语句的作用域问题

func main() {
	a := 10
	if a := 1; a > 0 {
		fmt.Println(a)
	}
	fmt.Println(a)
}
