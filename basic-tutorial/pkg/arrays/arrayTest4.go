package main

import "fmt"

// 不建议使用数组作为函数的参数
func terraform(planets [8]string) {
	for i := range planets {
		planets[i] = "New " + planets[i]
	}
}

func main() {
	planets := [...]string{
		"Mercury",
		"Venus",
		"Earth",
		"Mars",
		"Jupiter",
		"Saturn",
		"Uranus",
		"Neptune",
	}

	// 不会被改变
	terraform(planets)
	fmt.Println(planets)
}
