package basic

import "fmt"

func SliceExample() {
	s := []int{1, 2, 3, 4, 5}
	s2 := []int{1, 5: 2, 10: 3}
	s3 := make([]int, 3)
	fmt.Println(s)
	fmt.Println(s2)
	fmt.Println(s3)

	var slice = []int{1, 2, 3}
	for i := 0; i < len(slice); i++ {
		slice[i] += 10
	}
	for i, v := range slice {
		slice[i] = v * 2
	}
	fmt.Println(slice)

	slice = append(slice, 11, 12, 13, 14, 15)
	fmt.Println(slice)

}
