package basic

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFibonacci1(t *testing.T) {
	ast := assert.New(t)

	ast.Equal(0, fibonacci1(-1), "fibonacci1(-1) should be 0")
	ast.Equal(0, fibonacci1(0), "fibonacci1(0) should be 0")
	ast.Equal(1, fibonacci1(1), "fibonacci1(1) should be 1")
	ast.Equal(2, fibonacci1(3), "fibonacci1(3) should be 2")
	ast.Equal(233, fibonacci1(13), "fibonacci1(13) should be 233")
}

func TestFibonacci2(t *testing.T) {
	ast := assert.New(t)

	ast.Equal(0, fibonacci2(-1), "fibonacci2(-1) should be 0")
	ast.Equal(0, fibonacci2(0), "fibonacci2(0) should be 0")
	ast.Equal(1, fibonacci2(1), "fibonacci2(1) should be 1")
	ast.Equal(2, fibonacci2(3), "fibonacci2(3) should be 2")
	ast.Equal(233, fibonacci2(13), "fibonacci2(13) should be 233")
}

func BenchmarkFibonacci1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fibonacci1(20)
	}
}

func BenchmarkFibonacci2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fibonacci2(20)
	}
}
