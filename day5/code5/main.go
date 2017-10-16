// code5 project main.go
package main

//方法method的用法
import (
	"fmt"
)

type A struct{}
type TZ int

func main() {
	a := A{}
	b := a.Method(2)
	fmt.Println(b)
	fmt.Println("=========")

	var t TZ
	t.Print()       // Method Value
	(*TZ).Print(&t) // Method Expression
	fmt.Println(t)
	//两种方法都可以调用到Print
}

func (t *TZ) Print() {
	*t = 5
	fmt.Println("aaa")
}

func (a A) Method(x int) int {
	x = 5
	fmt.Println(a, x)
	return x
}

/*
func (a A) Method() {
	fmt.Println("aaa")
}

注意： 这样写是错误的，因为go语言中没有方法的重载！
*/
