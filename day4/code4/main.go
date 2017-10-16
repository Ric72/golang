// code4 project main.go
package main

//Go语言版冒泡排序
import (
	"fmt"
)

func main() {
	a := []int{1, 3, 21, 2, 5}
	fmt.Println(a)
	l := len(a)
	for i := 0; i < l; i++ {
		for j := i + 1; j < l; j++ {
			if a[i] > a[j] {
				temp := a[i]
				a[i] = a[j]
				a[j] = temp
			}
		}
	}
	fmt.Println(a)
}
