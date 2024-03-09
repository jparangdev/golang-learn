package basic

import "testing"

func TestSquared1(t *testing.T) {
	rst := squareTest(9)
	if rst != 81 {
		t.Errorf("fail test!")
	}
}
