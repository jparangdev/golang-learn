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
	fmt.Println(cap(slice))

	slice1 := make([]int, 3, 5)
	slice2 := append(slice1, 4, 5)
	fmt.Println("slice1:", slice1, len(slice1), cap(slice1))
	fmt.Println("slice2:", slice2, len(slice2), cap(slice2))

	slice1[1] = 100
	fmt.Println("after change")
	fmt.Println("slice1:", slice1, len(slice1), cap(slice1))
	fmt.Println("slice2:", slice2, len(slice2), cap(slice2))

	slice1 = append(slice1, 500)
	fmt.Println("after add")
	fmt.Println("slice1:", slice1, len(slice1), cap(slice1))
	fmt.Println("slice2:", slice2, len(slice2), cap(slice2))

	slice3 := append(slice1, 99, 98, 97)
	slice1[1] = 200
	fmt.Println("after over cap")
	fmt.Println("slice1:", slice1, len(slice1), cap(slice1))
	fmt.Println("slice2:", slice2, len(slice2), cap(slice2))
	fmt.Println("slice3:", slice3, len(slice3), cap(slice3))

}
