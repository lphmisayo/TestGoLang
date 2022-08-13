package main

import "fmt"

const (
	BIG   = 1 << 100
	SMALL = BIG >> 99
)

func needInt(x int) int {
	return x*10 + 1
}

func needFloat(x float64) float64 {
	return x * 0.1
}

func main() {
	var i float64 = 1.002
	k := i

	var x int = 3
	var j float64 = float64(x)

	fmt.Println(needInt(SMALL))
	fmt.Println(needFloat(BIG))
	fmt.Printf("Hello World%T%T", k, j)
}
