// code12 project main.go
package main

//闭包的应用

import (
	"fmt"
)

func main() {
	//此时f是一个匿名函数 func(int) int， 且外层函数x的值为5
	f := A(5)
	//再利用f(y int)进行对y的赋值
	fmt.Println(f(8))
	fmt.Println(f(9))
}

//将一个函数当作返回值的类型

func A(x int) func(int) int {
	fmt.Printf("%p\n", &x)
	return func(y int) int {
		fmt.Printf("%p\n", &x)
		return x + y
	}
}

//此时x一直是同一个内存地址上的值，不是值的拷贝，而是内存地址的引用
//所以x的内存地址一直都是相同的
