// code14 project main.go
package main

import (
	"fmt"
)

//根据for循坏的执行顺序判断出结果为3个3

func main() {
	for i := 0; i < 3; i++ {
		//defer func() {
		defer fmt.Println(i)
		//}()
	}
}
