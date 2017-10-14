// code5 project main.go
package main

import (
	"fmt"
)

var sum int

func main() {

	for n := 1; n < 10; n++ {
		sum += n

	}
	fmt.Println(sum)
}
