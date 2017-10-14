// code9 project main.go
package main

import (
	"fmt"
)

var array = []int{1, 5, 2, 4, 7, 8, 4}

func main() {
	for i := 0; i < len(array); i++ {
		fmt.Println(array[i])
	}

}
