package main

import "fmt"

type person struct {
	name, superpower string
	age              int
}

func birthday(p *person) {
	p.age++
}
func main() {
	rebecca := person{
		name:       "rebecca",
		superpower: "imagination",
		age:        14,
	}
	birthday(&rebecca)

	fmt.Printf("%+v\n", rebecca)
}
