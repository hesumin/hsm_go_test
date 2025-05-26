package main

import "fmt"

func main() {

	var a, b int
	var pwd int = 20221020

	// 用户输入
	fmt.Scan(&a)

	if a == pwd {
		fmt.Println("请再次输入密码：")
		fmt.Scan(&b)
		if b == pwd {
			fmt.Println("登录成功")
		} else {
			fmt.Println("密码错误")
		}

	} else {
		fmt.Println("登录失败")
	}
}
