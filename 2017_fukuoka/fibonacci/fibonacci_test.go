package fibonacci

import (
	"testing"
)

func TestFib(t *testing.T) {
	var result int

	result = fib(0)
	if result != 1 {
		t.Errorf("expect fib(0) to be 1, but got %d", result)
	}

	result = fib(1)
	if result != 1 {
		t.Errorf("expect fib(1) to be 1, but got %d", result)
	}

	result = fib(5)
	if result != 5 {
		t.Errorf("expect fib(5) to be 5, but got %d", result)
	}

	result = fib(10)
	if result != 55 {
		t.Errorf("expect fib(10) to be 55, but got %d", result)
	}

	result = fib(50)
	if result != 12586269025 {
		t.Errorf("expect fib(50) to be 12586269025, but got %d", result)
	}
}
