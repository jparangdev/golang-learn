package basic

import "fmt"

type Product struct {
	Name  string
	Price int
}

func MapExample() {
	m := make(map[int]Product)

	m[16] = Product{"pen", 500}
	m[46] = Product{"book", 5000}
	m[22] = Product{"pencil", 200}
	m[56] = Product{"eraser", 100}
	m[85] = Product{"ruler", 300}

	// Printing all elements in the map
	for _, product := range m {
		fmt.Println(product)
	}

	delete(m, 16)
	delete(m, 1)
	fmt.Println(m[16])
	fmt.Println(m[46])

}

/*
Remember: Big O notation describes the performance or complexity of an operation in the worst-case scenario, and it's used to describe how an algorithm scales as the amount of data involved increases.
Array and Slices:
	Accessing an element (e.g array[i]): O(1) - constant time complexity.
	Inserting or removing an element: O(n) - linear time complexity as all the other elements need to be shifted.
Lists (container/list):
	Inserting an element: O(1) - constant time complexity.
	Removing an element: O(1) - constant time complexity.
	Accessing an element: O(n) - linear time complexity as it may have to traverse the whole list.
Maps:
	Inserting an element: O(1) - constant time complexity (average), but could degrade to O(n) in the worst-case scenario.
	Removing an element: O(1) - constant time complexity (average), but could degrade to O(n) in the worst-case scenario.
	Accessing an element: O(1) - constant time complexity (average), but could degrade to O(n) in the worst-case scenario.
Of course, the actual performance can further depend on factors such as memory layout, cache levels, the cost of hash functions, speed of allocating and freeing memory, and so on, but Big O gives a useful abstracted idea of how they behave.
*/
