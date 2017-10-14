// code6 project main.go
package main

import "fmt"

func main() {
	var sum int

	for sum = 1; sum < 5; sum++ {

		sum += sum
		fmt.Println(sum)
	}

}
