// code2 project main.go
package main

import (
	"fmt"
)

func calc(a int) (int, int) {
	b := (a - 10) / 2
	c := b + 10
	return b, c
}

func main() {
	fmt.Println(calc(20))
}
