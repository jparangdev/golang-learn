package main

import (
	"fmt"
	"golang-learn/basic"
)

func main() {
	fmt.Println("======================================")
	basic.VariableExample()
	fmt.Println("======================================")
	basic.FmtExample()
	fmt.Println("======================================")
	basic.OperatorExample()
	fmt.Println("======================================")
	fmt.Println(basic.Add(3, 6))
	basic.PrintAvgScore("김일등", 80, 74, 95)
	basic.PrintAvgScore("송이등", 88, 92, 53)
	basic.PrintAvgScore("박삼등", 78, 73, 73)
	c, success := basic.Divide(9, 3)
	fmt.Println(c, success)
	d, success := basic.Divide(9, 0)
	fmt.Println(d, success)
	basic.PrintNo(4)
	fmt.Println("======================================")
	basic.ConstExample(2)
	fmt.Println("======================================")
	basic.ForExample()
	fmt.Println("======================================")
	basic.ArrayExample()
	fmt.Println("======================================")
	basic.StructExample()
	fmt.Println("======================================")
	basic.PointerExample()
	fmt.Println("======================================")
	basic.StringsExample()
	fmt.Println("======================================")
	basic.PackageExample()
	fmt.Println("======================================")
}
