package main

import "fmt"

func main() {
	var planets [8]string
	planets[0] = "Mercury"
	planets[1] = "Venus"
	planets[2] = "Farth"

	earth := planets[2]

	fmt.Println(earth)

	//打印数组长度
	fmt.Println(len(planets))
	// 判断是否为空
	fmt.Println(planets[3] == "")
}
