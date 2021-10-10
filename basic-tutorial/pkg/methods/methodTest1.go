package main

import "fmt"

func main() {
	//申明新类型
	type celsius float64

	const degrees = 20
	var temperature celsius = degrees

	temperature += 10

	fmt.Print(temperature)

}
