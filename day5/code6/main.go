// code6 project main.go
package main

import (
	"fmt"
)

// a:=0,调用 a.Increace()方法之后a的值就增加100

type A int

var a A

func main() {
	a = 0
	a.Increace()
	a.Increace()
}

func (a *A) Increace() {
	*a += 100
	fmt.Println(*a)
}
