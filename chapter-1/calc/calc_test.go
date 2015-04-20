package calc

import "testing"

func TestAdd(t *testing.T) {
	var v int
	v = Add(15,10)
	if v != 25 {
		t.Error("Expected 25, got ", v)
	}
}
func TestSubtract(t *testing.T) {
	var v int
	v = Subtract(15,10)
	if v != 5 {
		t.Error("Expected 5, got ", v)
	}
}
