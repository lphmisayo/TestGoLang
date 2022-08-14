package main

import (
	"fmt"
)

func sqrt(x float64, y int) float64 {

	z := 1.0

	for i := 0; i < y; i++ {
		z -= (z*z - x) / (2 * z)
		fmt.Printf("func sqrt return %f\n", z)
	}

	return z
}

func testSwitch(x int) string {

	switch x {
	case 1:
		return "选择1"
	case 2:
		return "选择2"
	case 3:
		return "选择3"
	default:
		return "未匹配"
	}
}

func main() {

	//x := float64(20)
	//y := 10
	switchX := 30
	//sqrt(x, y)
	//fmt.Printf("math sqrt return %f", math.Sqrt(x))
	fmt.Println(testSwitch(switchX))

}
