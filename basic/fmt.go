package basic

import "fmt"

// FmtExample is a function that demonstrates various ways to format and print values to the console.
//
// Example usage:
//
//	FmtExample()
//
// This function prints the values of variables `a`, `b` and `f` using different formatting options.
// It demonstrates the use of fmt.Print, fmt.Println and fmt.Printf functions.
//
// Example output:
//
//	a: 10 b: 20
//	a: 10 b: 20 f: 32412351235.2342
//	a: 10 b: 20 f:32412351235.234200
func FmtExample() {
	var a int = 10
	var b int = 20
	var f float64 = 32412351235.2342

	fmt.Print("a:", a, "b:", b)
	fmt.Println("a:", a, "b:", b, "f:", f)
	fmt.Printf("a: %d b: %d f:%f\n", a, b, f)
}
