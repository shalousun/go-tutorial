package main

import "fmt"

//复合字面值初始化数组

func main() {
	var planets = [...]string{
		"Mercury",
		"Venus",
		"Earth",
		"Mars",
		"Jupiter",
		"Saturn",
		"Uranus",
		"Neptune",
	}
	fmt.Print(planets)
}
