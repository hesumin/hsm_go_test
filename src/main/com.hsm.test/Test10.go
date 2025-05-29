package main

import "fmt"

// 递归函数计算阶乘
func factorial(n int) int {
	// 基准条件
	if n == 0 {
		return 1
	}
	fmt.Printf("进入递归：factorial(%d)\n", n)
	// 递归条件
	return n * factorial(n-1)
}

func main() {
	fmt.Println(factorial(10)) // 输出: 120
}
