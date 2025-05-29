package main

import "fmt"

func main() {
	var siteMap map[string]string
	siteMap = make(map[string]string)

	siteMap["Google"] = "guge"
	siteMap["Runoob"] = "菜鸟教程"
	siteMap["Taobao"] = "淘宝"
	siteMap["Zhihu"] = "知乎"

	for site := range siteMap {
		fmt.Println("site:", site, "url:", siteMap[site])
	}

	name, ok := siteMap["Facebook"]
	if ok {
		fmt.Println("name:", name, "ok:", ok)
	} else {
		fmt.Println("111111111111")
	}
}
