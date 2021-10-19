package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// go语言中json要求struct中的字段首字母大写，并且采用驼峰来命名规范，
//因此实际使用中可能需要添加tag
func main() {
	type location struct {
		Lat  float64 `json:"latitude"`
		Long float64 `json:"longitude"`
	}
	curiosity := location{-4.5895, 137.4417}
	bytes, err := json.Marshal(curiosity)
	exitOnError2(err)
	fmt.Println(string(bytes))
}

// exitOnError prints any errors and exits.
func exitOnError2(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
