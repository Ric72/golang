// code7 project main.go
package main

import (
	"fmt"
)

//copy之后会不会也改变了原来数组的元素内容
func main() {
	a1 := []int{1, 2, 3, 4}
	a2 := []int{5, 6, 7, 8}
	s1 := a1[:3]
	s2 := a2[:2]
	copy(s1, s2)
	fmt.Println(s1)
	fmt.Println(a1)
	//答案是 改变了！
}
