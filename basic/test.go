package basic

import "fmt"

func squareTest(x int) int {
	return x * x
}

func TestExample() {
	fmt.Printf("9 * 9 = %d\n", squareTest(9))
}
