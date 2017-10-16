// code8 project main.go
package main

import (
	"fmt"
)

//创建map的时候一定要先初始化才可以使用
func main() {
	m1 := make(map[int]string)
	m1[1] = "OK"
	//通过ok的值来判断值是否存在
	a, ok := m1[1]
	fmt.Println(m1, ok, a)

	//当map的value还为一个map的时候，要用ok来判断是否对第二个map进行初始化
	m2 := make(map[int]map[int]string)
	a, ok = m2[1][1]
	fmt.Println(ok)
	if !ok {
		m2[1] = make(map[int]string)
	}
	m2[1][1] = "OK"
	fmt.Println(m2[1][1])
	a, ok = m2[1][1]
	fmt.Println(a, ok)

	fmt.Println("================")
	m3 := make(map[int]string)

	//map中的赋值，只能通过 m[1] = xxx的方式赋值，不能先将 m[1]赋值给b，再给b赋值这种形式，无法传递
	m3[1] = "OK"
	b := m3[1]
	//这样子也不行， 打印出来b为aaa，但m3[1]依旧为OK
	b = "aaa"
	fmt.Println(b)
	fmt.Println(m3[1])
}
