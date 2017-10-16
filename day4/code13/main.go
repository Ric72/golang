package main

import (
	"fmt"
)

func main() {
	var j int = 5
	a := func() func() {
		j = 5
		var i int = 10
		return func() {
			fmt.Printf("i,j:%d,%d\n", i, j)
		}
	}
	//将一个无需参数返回值为匿名函数的函数赋值给a()

	a()()
	j *= 2
	fmt.Println(j)
	// i = 80
	// i*=2这样是错的
	a()()
}
