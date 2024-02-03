package basic

import (
	"fmt"
	"sort"
)

type Student struct {
	Name string
	Age  int
}

func (s *Student) String() string {
	return fmt.Sprintf("Student Name is %s and age is %d", s.Name, s.Age)
}

type Students []Student

func (s Students) Len() int           { return len(s) }
func (s Students) Less(i, j int) bool { return s[i].Age < s[j].Age }
func (s Students) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

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

	fmt.Println("Remove an element at a index")
	slice4 := []int{1, 2, 3, 4, 5, 6}
	idx := 2
	slice5 := append(slice4[:idx], slice4[idx+1:]...)
	fmt.Println(slice5)

	fmt.Println("Add an element at a index")
	slice6 := append(slice5[:idx], append([]int{100}, slice5[idx:]...)...)
	fmt.Println(slice6)

	fmt.Println("Sort Ints")
	sort.Ints(slice6)
	fmt.Println(slice6)

	s := []Student{
		{Name: "John", Age: 18},
		{Name: "Jane", Age: 20},
		{Name: "Mark", Age: 19},
		{Name: "Lisa", Age: 17},
		{Name: "Bob", Age: 21},
	}

	sort.Sort(Students(s))

	fmt.Println("Sorted Students by Age")
	for _, student := range s {
		fmt.Println(student.Name, student.Age)
	}
}
