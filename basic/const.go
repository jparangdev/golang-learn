package basic

import "fmt"

//const Java int = 1
//const Golang int = 2
//const JavaScript int = 3

const (
	Java int = iota + 1
	Golang
	JavaScript
)

const Python = 4

func ConstExample(language int) {
	if language == Java {
		fmt.Println("Spring boot")
	} else if language == Golang {
		fmt.Println("Gin")
	} else if language == JavaScript {
		fmt.Println("Node")
	} else if language == Python {
		fmt.Println("Django")
	} else {
		fmt.Println("??")
	}
}
