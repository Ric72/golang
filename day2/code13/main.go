// code13 project main.go
package main

import "fmt"

func main() {
	var a []int
	printSlice("a", a)

	// the slice grows as needed.
	a = append(a, 1, 8, 8, 8, 8)
	printSlice("a", a)

}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}
