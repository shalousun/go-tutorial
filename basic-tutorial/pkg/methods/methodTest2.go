package main

import "fmt"

type celsius float64
type kelvin float64

func kelvinToVCelsius(k kelvin) celsius {
	return celsius(k - 273.15)
}

// 方法和kelvin类型关联，kelvin类型可以直接调用
func (k kelvin) celsius() celsius {
	return celsius(k - 273.15)
}

func main() {
	var k kelvin = 294.0
	var c celsius

	c = kelvinToVCelsius(k)
	c = k.celsius()
	fmt.Print(c)
}
