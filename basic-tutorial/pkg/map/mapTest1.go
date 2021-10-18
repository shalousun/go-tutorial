package main

import "fmt"

func main() {
	temperature := map[string]int{
		"Earth": 15,
		"Mars":  -65,
	}
	temp := temperature["Earth"] // get value from map
	fmt.Println("On overage the Earth is %v ", temp)

	temperature["Earth"] = 16
	temperature["Venus"] = 464

	fmt.Println(temperature)

	moon := temperature["Moon"] // 不存在的key

	fmt.Println(moon)

	if moon, ok := temperature["Moon"]; ok {
		fmt.Println("On overage the Earth is %v ", moon)
	} else {
		fmt.Println("Where is the moon?")
	}
}
