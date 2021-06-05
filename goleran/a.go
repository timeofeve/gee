package main

import (
	"fmt" 
)
func main() {
	type location struct {
		lat, long float64
	}
	curiosity := location{-4.11, 1221.33}
	fmt.Printf("%+v\n", curiosity)
	fmt.Printf("%v\n", curiosity)
}
