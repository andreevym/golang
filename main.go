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

	simpleArray := [3]int32{}
	simpleArray[0] = 1
	simpleArray[0] = 2
	simpleArray[0] = 4
	fmt.Println(simpleArray)

	myArray := [...]int32{1, 2, 4, 16, 32, 64, 128, 256, 512, 1024, 2048, 4096}
	fmt.Println(myArray)

	firstSlice := myArray[:]
	firstSlice = append(firstSlice, 8194)
	fmt.Println(firstSlice)

	secondSlice := make([]int32, 0)
	secondSlice = append(secondSlice, 1)
	secondSlice = append(secondSlice, 2)
	secondSlice = append(secondSlice, 3)
	fmt.Print(secondSlice)

	thirdSlice := []int32{1, 2, 4, 16, 32, 64, 128, 256, 512, 1024, 2048, 4096}
	thirdSlice = append(thirdSlice, 1)
	thirdSlice = append(thirdSlice, 2)
	thirdSlice = append(thirdSlice, 3)
	fmt.Print(thirdSlice)
}
