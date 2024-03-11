package basic

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSquared1(t *testing.T) {
	rst := squareTest(9)
	if rst != 81 {
		t.Errorf("fail test!")
	}
}

func TestSquared2(t *testing.T) {
	ass := assert.New(t)
	ass.Equal(81, squareTest(9), "square(9) should be 81")
}

func TestSquared3(t *testing.T) {
	ass := assert.New(t)
	ass.Equal(9, squareTest(3), "square(3) should be 9")
}
