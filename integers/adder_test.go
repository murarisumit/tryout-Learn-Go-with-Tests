package integers

import "testing"

func TestAdder(t *testing.T) {
	sum := Adder(2, 4)
	expected := 6

	if sum != expected {

		t.Errorf("Expected '%d' but got '%d'", expected, sum)

	}
}
