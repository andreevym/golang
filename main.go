package main

import (
	"fmt"
)

const (
	main_title  = "let's Go"
	array_title = "try Array"
)

const (
	first = 2 * iota
	second
	third
)

func main() {
	fmt.Println(main_title)
	fmt.Println(third)
	tryArray()
}

func tryArray() {
	fmt.Println(array_title)
	array := [...]int32{1, 2, 4, 16, 32, 64, 128, 256, 512, 1024, 2048, 4096}
	fmt.Println(array)
}
