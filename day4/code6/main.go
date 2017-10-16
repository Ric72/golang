// code6 project main.go
package main

import (
	"fmt"
)

//让slice指向一个完整的数组
func main() {
	a := [...]int{1, 3, 2, 3}
	s1 := a[:]
	fmt.Printf("%p\n", s1)
	fmt.Printf("%v\n", a)
}
