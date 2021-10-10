package main

import (
	"fmt"
	"strings"
)

// 切片作为方法的参数，数组是有长度的
func hyperspace(worlds []string) {
	for i := range worlds {
		worlds[i] = strings.TrimSpace(worlds[i])
	}
}
func main() {
	planets := []string{" Venus   ", "Earth  ", " Mars"}

	hyperspace(planets)

	fmt.Println(strings.Join(planets, ""))
}
