package main

import (
	"fmt"
)

func main() {
	m := map[string]int{
		"apple":  1,
		"banana": 2,
		"orange": 3,
	}

	n := map[string]int{
		"apple":  1,
		"banana": 2,
		"pear":   3,
	}
	fmt.Println(m["apple"], n["banana"])
	// for k,vm:=range m{
	// 	if
	// }

}
