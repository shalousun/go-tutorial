package main

import "fmt"

func main() {
	neptune := "Neptune"
	// 对字符串进行切片
	tune := neptune[3:]

	fmt.Println(tune)

	neptune = "Poseidon"
	// 不会改变切片
	fmt.Println(tune)
}
