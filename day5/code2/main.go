// code2 project main.go
package main

import (
	"fmt"
)

//利用recover机制，保证函数在遇到错误之后还能继续往下执行

func main() {
	A()
	B()
}

func A() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("遇到错误，但是被recover了，函数得以继续执行")
		}
	}()
	panic("aaa")
}

func B() {
	fmt.Println("Func B")
}
