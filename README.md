## golang
sample with golang

## How to run
go run main.go

## Constant
pay attention on and don't use curly braces

const (
	main_title  = "let's Go"
)

## Array
can by with fixed size or scalable

## How to create "Simple array with fixed size" ?
	array := [...]int32{1, 2, 4, 16, 32, 64, 128, 256, 512, 1024, 2048, 4096}

## How to create "Slice - scalable array" ?

### First slice
myArray := [...]int32{1, 2, 4, 16, 32, 64, 128, 256, 512, 1024, 2048, 4096}
fmt.Println(myArray)

firstSlice := myArray[:]
firstSlice = append(firstSlice, 8194)
fmt.Println(firstSlice)

### Second slice
secondSlice := make([]int32, 0)
secondSlice = append(secondSlice, 1)
secondSlice = append(secondSlice, 2)
secondSlice = append(secondSlice, 3)
fmt.Print(secondSlice)

### Third slice
thirdSlice := []int32{1, 2, 4, 16, 32, 64, 128, 256, 512, 1024, 2048, 4096}
thirdSlice = append(thirdSlice, 1)
thirdSlice = append(thirdSlice, 2)
thirdSlice = append(thirdSlice, 3)
fmt.Print(thirdSlice)