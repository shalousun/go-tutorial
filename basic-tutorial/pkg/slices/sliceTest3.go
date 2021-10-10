package main

import "fmt"

func main() {
	dwarfArray := [...]string{"Ceres", "Pluto", "Haumea", "Makemake", "Eris"}

	// 使用数组申明一个slice
	dwarfSlice := dwarfArray[:]

	// 另一种方式申明slice
	dwarfs := []string{"Ceres", "Pluto", "Haumea", "Makemake", "Eris"}

	fmt.Println(dwarfSlice, dwarfs)
}
