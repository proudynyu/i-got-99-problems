package main

import (
	"example.com/99_problems/src"
	"fmt"
)

func main() {
	list := []string{"a", "a", "a", "a", "b", "c", "c", "a", "a", "d", "d", "e", "e", "e", "e"}
	test := src.PackConsecutivesDuplicates(list)
	fmt.Println(test)
}
