package main

import "fmt"

// 获取slice的长度和容量
func dump2(label string, slice []string) {
	fmt.Println("%v: length %v, capacity %v %v\n", label, len(slice), cap(slice), slice)
}
func main() {
	planets := []string{
		"Mercury", "Venus", "Earth", "Mars",
		"Jupiter", "Saturn", "Uranus", "Neptune",
	}
	terrestrial := planets[0:4:4] //通过第三个索引指定容量
	worlds := append(terrestrial, "Ceres")

	dump2("planets", planets)
	dump2("terrestrial", terrestrial)
	dump2("worlds", worlds)
}
