// code9 project main.go
package main

//利用slice来对map进行一个排序操作
//因为map存取值的时候是无序的
import (
	"fmt"
	"sort"
)

func main() {
	m := map[int]string{1: "a", 2: "b", 3: "c", 4: "d", 5: "e"}
	s := make([]int, len(m))
	i := 0
	for k, _ := range m {
		s[i] = k
		i++
	}
	sort.Ints(s)
	for _, i = range s {
		fmt.Println(m[i])
	}

}
