// code10 project main.go
package main

//将一个map的键值对进行交换
import (
	"fmt"
)

func main() {
	m1 := make(map[int]string)
	a := 'a'
	for i := 0; i < 10; i++ {
		m1[i] = string(a)
		a++
	}
	m2 := make(map[string]int)
	for k, v := range m1 {
		m2[v] = k
	}
	fmt.Println("m1", m1)
	fmt.Println("m2", m2)
}
