package basic

import "fmt"

// OperatorExample demonstrates the usage of various operators in Go.
func OperatorExample() {

	var x int32 = 7
	var y int32 = 3

	var s float32 = 3.14
	var t float32 = 5

	fmt.Println("x + y = ", x+y)
	fmt.Println("x - y = ", x-y)
	fmt.Println("x * y = ", x*y)
	fmt.Println("x / y = ", x/y)
	fmt.Println("x % y = ", x%y)

	fmt.Println("s * t = ", s*t)
	fmt.Println("s / t = ", s/t)

	var x1 int8 = 34   // 00100010
	var x2 int16 = 34  // 00000000 00100010
	var x3 uint8 = 34  // 00100010
	var x4 uint16 = 34 // 00000000 00100010

	fmt.Printf("^%d = %5d, \t %08b\n", x1, ^x1, uint8(^x1))
	fmt.Printf("^%d = %5d, \t %16b\n", x2, ^x2, uint16(^x2))
	fmt.Printf("^%d = %5d, \t %08b\n", x3, ^x3, ^x3)
	fmt.Printf("^%d = %5d, \t %16b\n", x4, ^x4, ^x4)
}
