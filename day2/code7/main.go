// code7 project main.go
package main

import (
	"fmt"
)

type Product struct {
	name        string
	price       float64
	discription string
}

func main() {
	p1 := Product{"basket", 12.5, "还不错"}
	p2 := Product{"basketall", 125, "还不错2"}
	fmt.Println(p1.name, p2.discription)
}
