package main

import (
	"fmt" 
)
func main() {
	t := []float64{
		1.0, 2.0, 3.0, 4.0, 5.0, 1.0, 2.0, 3.0,
	}
	f := make(map[float64]int)

	for _, t := range t {
		f[t]++
		fmt.Println(t)
	}
	for t, num := range f {
		fmt.Printf("%+.2f dd %d times\n", t, num)
	}
}
