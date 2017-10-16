// code4 project main.go
package main

import (
	"fmt"
)

//匿名字段名的默认名称与外层的字段名重合的时候怎么办
//就不能使用匿名字段，必须给结构体命名

type human struct {
	Sex int
}

type teacher struct {
	h     human
	human string
	age   int
}

func main() {
	a := teacher{human: "abc", age: 19, h: human{Sex: 1}}
	a.h.Sex = 10
	fmt.Println(a)
}
