package basic

import (
	"fmt"
	"os"
)

func Function2Example() {
	fmt.Println(sum(1, 2, 3, 4, 5))
	fmt.Println(sum(6, 7, 8))

	defferExample()
}

func sum(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}

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
