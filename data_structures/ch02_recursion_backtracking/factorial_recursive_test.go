package rb

import "testing"

func TestFactorial(t *testing.T) {
	if factorial(5) != 120 {
		t.Error(factorial(5))
	}

	if common(6) != 720 {
		t.Error(common(6))
	}
}