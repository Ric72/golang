// code3 project main.go
package main

import (
	"fmt"
)

func main() {

	str1 := "a b cd e fghi jasdkljas "
	fmt.Println(str1[3])
	fmt.Println(str1[1:5])
	fmt.Println(str1[:5])
	fmt.Println(str1[10:])
	fmt.Println(len(str1))
	fmt.Println(len([]rune(str1)))

}
