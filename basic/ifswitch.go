package basic

import "fmt"

func ForExample() {
	temp := 33

	if temp > 28 {
		fmt.Println("turn on the air conditioner!!")
	} else if temp <= 3 {
		fmt.Println("turn on the heater")
	} else {
		fmt.Println("zZzZzZ")
	}

	price := 35000
	if price > 50000 {
		if HasRichFriend() {
			fmt.Println("I lost my wallet :) ")
		} else {
			fmt.Println("Let's split the bill.")
		}
	} else if price >= 30000 && price <= 50000 {
		if GetFriendCount() > 3 {
			fmt.Println("where is my wallet?")
		} else {
			fmt.Println("Let's split the bill ..")
		}
	} else {
		fmt.Println("I'll pay for it !!")
	}

	if age, ok := getMaxAge(); ok && age < 20 {
		fmt.Println("you are so young", age)
	} else if ok && age < 30 {
		fmt.Println("nice!", age)
	} else if ok {
		fmt.Println("beautiful!")
	} else {
		fmt.Println("Error!")
	}

	switch age, ok := getMaxAge(); true {
	case age < 10 && ok:
		fmt.Println("child")
	case age < 20 && ok:
		fmt.Println("Teenager")
	case age < 30 && ok:
		fmt.Println("20s")
		break
	case age < 40 && ok:
		fmt.Println("30s")
		fallthrough
	default:
		fmt.Println("my age is", age)
	}
}

func HasRichFriend() bool {
	return true
}

func GetFriendCount() int {
	return 3
}
func getMaxAge() (int, bool) {
	return 33, true
}
