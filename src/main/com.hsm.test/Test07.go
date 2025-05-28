package main

import "fmt"

func main() {
	var munbers []int

	printSlice7(munbers)

	if munbers == nil {
		fmt.Printf("munbers is nil\n")
	}
}

func printSlice7(x []int) {
	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)

}
