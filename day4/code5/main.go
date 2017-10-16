// code5 project main.go
package main

import (
	"fmt"
)

//slice中的append对内存地址的影响
func main() {
	s1 := make([]int, 5)
	s2 := s1[1:2]
	//由于此时append之后已经超过最大容量，所以再改变元素时，已经不是在同一个内存地址，所以不会对s1造成影响
	//s2 = append(s2, 3, 3, 21, 12, 3, 5)

	//但此时只添加一个元素，还在最大容量范围之内，所有对s2的操作都与s1在同一个内存地址，所以s1也会被影响
	s2 = append(s2, 3)
	s2[0] = 1
	fmt.Println(s2)
	fmt.Println(s1)
}
