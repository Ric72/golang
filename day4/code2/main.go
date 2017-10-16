// code2 project main.go
package main

//标签的使用
import (
	"fmt"
)

func main() {
LABEL1:
	for i := 0; i < 1; i++ {
		for {
			fmt.Println(i)
			continue LABEL1

		}
	}
	fmt.Println("Hello World!")
}
