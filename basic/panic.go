package basic

import "fmt"

// PanicExample is a function that demonstrates the usage of panic and recover in Go.
//
// The function calls another function named f, which internally calls a function named g.
// If an error occurs in g, it will panic and the panic will be recovered in f using a defer statement.
// After recovering from the panic, the program continues running and prints "program is running".
func PanicExample() {
	f()
	fmt.Println("program is running")
}

// f is a function that demonstrates the usage of panic and recover in Go.
//
// The function starts by printing "Function F is start".
// It then uses a defer statement to recover from any panics that may occur.
// If a panic occurs, the recovered value is printed as "Recovered from panic: <recovered value>".
//
// After recovering from the panic, the function calls the function g.
// Finally, "F end" is printed to indicate the end of the function.
//
//	program is running
func f() {
	fmt.Println("Function F is start")
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()

	g()
	fmt.Println("F end")
}

// g is a function that calls the function h with different arguments to perform division operations.
// It prints the results of these operations using formatted output. If h encounters
func g() {
	fmt.Printf("9 / 3 = %d\n", h(9, 3))
	fmt.Printf("9 / 0 = %d\n", h(9, 0))
}

// h performs integer division between two integers.
// If the second integer is zero, it will panic with the message "division by zero".
func h(a, b int) int {
	if b == 0 {
		panic("division by zero")
	}
	return a / b
}
