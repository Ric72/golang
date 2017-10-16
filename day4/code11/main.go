// code11 project main.go
package main

import (
	"fmt"
)

//函数的一些用法
func main() {
	f := func() func() {
		return func() {
			fmt.Println("aaa")
		}
	}()
	f() //打印出OKK，这个就是匿名函数的用法
}
