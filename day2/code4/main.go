// code4 project main.go
package main

import (
	"fmt"
)

const (
	Big   = 2
	Small = 5
)

func main() {
	var res = Big << Small
	fmt.Println(res)
	fmt.Println(Small << Big)
	fmt.Println(float64(Small >> 3))
}
