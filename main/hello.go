package main

import (
	"fmt"
	"math"
)

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

func testFor(x int, y int) int {

	sum := 0

	for i := x; i < y; i++ {
		sum++
	}

	return sum
}

func testWhileInGo(x int) int {

	sum := 0

	for sum < x {
		sum++
	}

	return sum
}

func testIf(x int) bool {

	if x == 0 {
		return false
	} else {
		return true
	}

}

func testIfMax(x, y float64) float64 {

	max := 10.0

	if v := math.Pow(x, y); v < max {
		return v
	}

	return max

}

func main() {
	//var i float64 = 1.002
	//k := i

	//var x int = 3
	//var j float64 = float64(x)

	//fmt.Println(needInt(SMALL))
	//fmt.Println(needFloat(BIG))
	//fmt.Printf("Hello World%T%T", k, j)

	fmt.Printf("TESTFOR RETURN %d\n", testFor(4, 20))
	fmt.Printf("testWhileInGo RETURN %d\n", testWhileInGo(20))
	fmt.Println(testIf(1))
	fmt.Println(testIfMax(3, 3))
}
