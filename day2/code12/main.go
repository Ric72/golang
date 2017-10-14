// code12 project main.go
package main

import (
	"fmt"
)

func main() {
	var a = [4]int{1, 2, 3, 5}
	var b = []int{1, 2, 3}
	var c []int

	s1 := a[1:3]
	s2 := make([]int, 7, 10)

	s1 = append(s1, 5, 7, 8, 8, 8)
	s2 = append(s2, 5, 7, 7, 7, 7, 6, 5, 4, 3, 7, 5, 4, 3, 5, 34, 4, 5)

	s3 := append(c, 3, 2, 3, 8, 8, 8, 8, 8, 8)

	fmt.Println(len(a), cap(a), a)
	fmt.Println(len(b), cap(b), b)
	fmt.Println(len(s1), cap(s1), s1)
	fmt.Println(len(s2), cap(s2), s2)
	fmt.Println(len(c), cap(c), c)
	fmt.Println(len(s3), cap(s3), s3)

}
