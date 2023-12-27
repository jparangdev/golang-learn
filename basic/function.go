package basic

import "fmt"

// Add takes two integers and returns their sum
func Add(a int, b int) int {
	return a + b
}

// PrintAvgScore writes the average score of a student to the console.
//
// Parameters:
// - name: the name of the student (string)
// - math: the math score of the student (int)
// - eng: the English score of the student (int)
// - history: the history score of the student (int)
//
// Example usage:
//
//	PrintAvgScore("김일등", 80, 74, 95)  // 김일등 님 평균 점수는 83 입니다.
//	PrintAvgScore("송이등", 88, 92, 53)  // 송이등 님 평균 점수는 77 입니다.
//	PrintAvgScore("박삼등", 78, 73, 73)  // 박삼등 님 평균 점수는 74 입니다.
func PrintAvgScore(name string, math int, eng int, history int) {
	total := math + eng + history
	avg := total / 3
	fmt.Println(name, "님 평균 점수는", avg, "입니다.")
}

// Divide divides two integers and returns the result and success status.
// If the second integer is zero, the result will be 0 and success will be false.
// Otherwise, the result will be the division of the first integer by the second integer and success will be true.
//
// Usage Example:
//
//	result, success := Divide(9, 3)
//	fmt.Println(result, success)
//
//	result, success := Divide(9, 0)
//	fmt.Println(result, success)
//
// Output:
//
//	3 true
//	0 false
func Divide(a, b int) (result int, success bool) {
	if b == 0 {
		result = 0
		success = false
		return
	}
	result = a / b
	success = true
	return
}

func PrintNo(n int) {
	if n == 0 {
		return
	}
	fmt.Println(n)
	PrintNo(n - 1)
	fmt.Println("After", n)
}
