package basic

import "fmt"

func ArrayExample() {
	// Declaring and initializing an array 't' of float64 with 5 elements - 24.0, 25.9, 27.8, 26.9, 26.2.
	var t = [5]float64{24.0, 25.9, 27.8, 26.9, 26.2}

	// A loop that runs for the length of array 't' and prints each element.
	for i := 0; i < 5; i++ {
		fmt.Println(t[i])
	}

	// Declaring and initializing an array 'nums' of integers. Using [...], the compiler automatically counts the number of elements.
	// The initial values are - 10, 20, 30, 40, 50. Later the value at index 2 in array 'nums' is changed to 300.
	nums := [...]int{10, 20, 30, 40, 50}
	nums[2] = 300

	// A range loop that iterates over the array 't' and prints the index 'i' and the corresponding element 'v'.
	for i, v := range t {
		fmt.Println(i, v)
	}

	// A range loop that iterates over the array 't' and prints the element 'v'. Does not care about index so it's ignored with '_'.
	for _, v := range t {
		fmt.Println(v)
	}

	// Declaring and initializing two new arrays 'a' and 'b' with integer elements.
	a := [5]int{1, 2, 3, 4, 5}
	b := [5]int{500, 400, 300, 200, 100}

	// Assigning the elements of array 'a' to array 'b'. The elements of 'b' are now identical to 'a'.
	b = a

	// Using a range loop to print out each index position 'i' and corresponding value 'v' of array 'b'.
	for i, v := range b {
		fmt.Println(i, v)
	}

	// Declaring and initializing a 2D array 'ab' that consists of 2 arrays, each containing 5 elements.
	ab := [2][5]int{
		{1, 2, 3, 4, 5},
		{6, 7, 8, 9, 10},
	}

	// A nested loop to iterate over the 2D array 'ab'. For each array within 'ab', the loop prints all indexes and corresponding values.
	for _, arr := range ab {
		for _, v := range arr {
			fmt.Println(v)
		}
	}
}
