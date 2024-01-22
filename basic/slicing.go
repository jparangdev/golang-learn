package basic

import "fmt"

func SlicingExample() {
	array := [5]int{1, 2, 3, 4, 5}
	slice := array[1:2]

	fmt.Println("array:", array)
	fmt.Println("slice:", slice, len(slice), cap(slice))

	array[1] = 100

	fmt.Println("After change second element")
	fmt.Println("array:", array)
	fmt.Println("slice:", slice, len(slice), cap(slice))

	slice = append(slice, 500)

	fmt.Println("After append to slice")
	fmt.Println("array:", array)
	fmt.Println("slice:", slice, len(slice), cap(slice))

	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := make([]int, len(slice1))
	for i, v := range slice1 {
		slice2[i] = v
	}

	slice1[1] = 100
	fmt.Println(slice1)
	fmt.Println(slice2)

	slice3 := make([]int, 10)
	cnt1 := copy(slice3, slice2)

	fmt.Println(cnt1, slice3)

	slice4 := []int{1, 2, 3, 4, 5, 6}
	idx := 2
	slice5 := append(slice4[:idx], slice4[idx+1:]...)
	fmt.Println(slice5)
}
