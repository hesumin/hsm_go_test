package main

import "fmt"

func main() {
	var score = 50

	switch score {
	case 90:
		fmt.Println("A")
	case 80:
		fmt.Println("B")
	case 60:
		fmt.Println("C")
	default:
		fmt.Println("D")
	}
}
