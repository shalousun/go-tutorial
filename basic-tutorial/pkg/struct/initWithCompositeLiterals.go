package main

import "fmt"

// 通过字面组合初始化
// type申明结构体可以复用
type location1 struct {
	lat  float64
	long float64
}

func main() {
	opportunity := location1{
		lat:  -1.9462,
		long: 354.4734,
	}
	fmt.Println(opportunity)

	insight := location1{
		lat:  4.5,
		long: 135.9,
	}
	fmt.Println(insight)

	fmt.Printf("%v\n", insight)
	// 打印带上字段
	fmt.Printf("%+v\n", insight)
}
