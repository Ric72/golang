// code4 project main.go
package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	m := make(map[string]int) // map对象
	c := strings.Fields(s)    // []byte
	for _, v := range c {
		m[v] += 1 //没有k，v就添加  有的话就修改v
	}
	return m
}
func main() {
	wc.Test(WordCount)
}
