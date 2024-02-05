package basic

import (
	"fmt"
	"os"
)

func Function2Example() {
	fmt.Println(sum(1, 2, 3, 4, 5))
	fmt.Println(sum(6, 7, 8))

	defferExample()
	fn := getOperator("+")
	fmt.Println(fn(1, 2))
	fn = getOperator("*")
	fmt.Println(fn(3, 4))
	fn = getOperator("%")
	fmt.Println(fn)
	CaptureLoop1()
	CaptureLoop2()
}

/*
This function, sum, is an example of how to define a function that can receive a variable number of arguments in Go.
The nums ...int parameter allows the function to accept any number of int arguments.
*/
func sum(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}

/*
The defer keyword in Go is used to schedule a function call (the deferred function) to be run immediately before the function that the defer statement is in returns.
It's typically used to perform cleanup tasks such as closing files, network connections, or database connections.
In the defferExample function, three deferred function calls are scheduled:
fmt.Println("first?") is the first deferred call,
and therefore will be the last one to be executed when defferExample returns.
f.Close() is scheduled second, and so will be execute just before fmt.Println("first?").
fmt.Println("last?") is the third deferred call, which means it will be the first one to execute among these three deferred calls when defferExample returns.
Essentially, defer calls are placed on a stack and then popped off in a LIFO (last-in-first-out) order when the function returns.
*/
func defferExample() {
	f, err := os.Create("deffer.txt")
	if err != nil {
		fmt.Println("fail to create a file")
		return
	}

	defer fmt.Println("last?")
	defer f.Close()
	defer fmt.Println("first?")

	fmt.Println("write Hello world on file")
	fmt.Fprintln(f, "Hello world")
}

/*
Using abbreviations, we can simplify the usage of functions as variables.
*/
type opFunc func(int, int) int

/*
this function is takes a string as an arguments and returns a function that performs the corresponding operation.
'opFunc' is an abbreviation for a type of function above
*/
func getOperator(op string) opFunc {
	switch op {
	case "+":
		return func(a, b int) int {
			return a + b
		}
	case "*":
		return func(a, b int) int {
			return a * b
		}
	default:
		return nil
	}
}

/*
The functions 'CaptureLoop1' and 'CaptureLoop2' demonstrate important aspects
of variable scope and lifetimes within closures and loops

In 'CaptureLoop1' function, we see that 'i' is in the scope of the for loop.
As a result, all the function literals created within the loop refer to the same instance of 'i'
this is a consequence of pointer copying, and all functions end up sharing the final value of 'i'

In contrast, in 'CaptureLoop2' each iteration creates a new instance of 'v',
so each function literal captures a separate 'v' instance. Due to this value copy, each function prints a distinct value
*/

func CaptureLoop1() {
	f := make([]func(), 3)
	fmt.Println("value loop")
	for i := 0; i < 3; i++ {
		f[i] = func() {
			fmt.Println(i)
		}
	}

	for i := 0; i < 3; i++ {
		f[i]()
	}
}

func CaptureLoop2() {
	f := make([]func(), 3)
	fmt.Println("value loop2")
	for i := 0; i < 3; i++ {
		v := i
		f[i] = func() {
			fmt.Println(v)
		}
	}

	for i := 0; i < 3; i++ {
		f[i]()
	}
}
