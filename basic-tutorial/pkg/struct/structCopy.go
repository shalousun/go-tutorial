package main

import "fmt"

func main() {
	type location struct {
		lat, long float64
	}

	bradbury := location{-4.5895, 137.4417}
	// struct会进行完整的复制，复制后互补影响
	curiosity := bradbury
	curiosity.long += 0.0106

	fmt.Println(bradbury, curiosity)
}
