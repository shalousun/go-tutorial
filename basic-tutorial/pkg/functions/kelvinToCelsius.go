package main

import "fmt"

// 函数测试
func kelvinToCelsius(K float64) float64 {
	K -= 237.15
	return K
}

func main() {
	kelvin := 294.0
	// 同包直接调用，函数按值传递参数
	celsius := kelvinToCelsius(kelvin)
	fmt.Print(kelvin, " K is", celsius, " C")
}
