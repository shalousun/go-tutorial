package main

import "fmt"

type Person struct {
	name, superpower string
	age              int
}

func (p *Person) birthday1() {
	p.age++
}

func main() {
	terry := &Person{
		name: "Terry",
		age:  15,
	}
	terry.birthday1()
	fmt.Printf("%+v\n", terry)

	nathan := Person{
		name: "Nathan",
		age:  17,
	}
	nathan.birthday1()
	fmt.Printf("%+v\n", nathan)
}
