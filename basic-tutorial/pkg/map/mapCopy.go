package main

import "fmt"

func main() {
	planets := map[string]string{
		"Earth": "Sector ZZ9",
		"Mars":  "Sector ZZ9",
	}

	planets2 := planets         //map不会被复制，执行同一块内存
	planets["Earth"] = "whoops" //修改了planets也会导致上面copy的被修改

	fmt.Println(planets)
	fmt.Println(planets2)

	delete(planets, "Earth")
	fmt.Println(planets2)
}
