package basic

import "fmt"

// VariableExample demonstrates the declaration and usage of various types of variables.
// It showcases the declaration and initialization of variables using different syntaxes such as
// var, short variable declaration, and type inference.
// It also demonstrates the conversion of variables between different types.
//
// Example Usage:
//
//	VariableExample()
//
// Output:
//
//	3 0 4 5
//	3.1415 hello golang! true
//	3.5 3 10.5 10
func VariableExample() {

	var a int = 3
	var b int
	var c = 4
	d := 5
	fmt.Println(a, b, c, d)

	var e = 3.1415
	f := "hello golang!"
	g := true
	fmt.Println(e, f, g)

	var h float64 = 3.5
	var i int = 3
	j := h * float64(i)
	k := int(h) * i // h를 int로 변환 (버림 발생)
	fmt.Println(h, i, j, k)
}
