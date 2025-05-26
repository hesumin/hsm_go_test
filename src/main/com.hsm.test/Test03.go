package main

import "fmt"

func main() {

	//sum := 0
	//for i := 1; i <= 10; i++ {
	//	sum += i
	//	fmt.Println(sum)
	//}

	//for i := 0; i < 5; i++ {
	//	for j := 0; j < 5; j++ {
	//		fmt.Print("* ")
	//	}
	//	fmt.Println()
	//}
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d * %d = %d\t", j, i, i*j)
		}
		fmt.Println()
	}
}
