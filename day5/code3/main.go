// code3 project main.go
package main

//分析结果

import (
	"fmt"
)

func main() {
	var fs = [4]func(){}

	for i := 0; i < 4; i++ {
		defer fmt.Println("defer i=", i)
		//	defer func(i int) { fmt.Println("defer_closure i=", i) }(i)
		defer func() { fmt.Println("defer_closure i=", i) }()
		//注意看上两句话的区别，如果是第一句话来运行的话，i虽然在闭包里面，但是是作为参数传递进去的，并不是对i的引用，而是对值的拷贝，此时的结果为 3 2 1 0
		//但是第二句话中，i不是作为参数传递进去的，所以是对i的引用，循环结束后，i的值为4，所以此时为 4 4 4 4

		fs[i] = func() {
			fmt.Println("closure i=", i)
		}
	}

	for _, f := range fs {
		f()
	}
}
