// code1 project main.go
package main

import (
	"fmt"
)

//panic的用法

func main() {
	A()
	B()
}

func A() {
	panic("此处发生错误")
}

func B() {
	fmt.Printf("Func B")
}

//由于在调用函数A的时候已经发生了错误，所以B不会执行
