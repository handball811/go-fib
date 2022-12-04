package internal

import "testing"

func TestFibGen(t *testing.T) {
	testCases := []struct {
		name   string
		n      int
		result int
		err    error
	}{
		{
			name:   "Fibonacci number of 0",
			n:      0,
			result: 0,
			err:    nil,
		},
		{
			name:   "Fibonacci number of 1",
			n:      1,
			result: 1,
			err:    nil,
		},
		{
			name:   "Fibonacci number of 5",
			n:      5,
			result: 5,
			err:    nil,
		},
		{
			name:   "Fibonacci number of 10",
			n:      10,
			result: 55,
			err:    nil,
		},
		{
			name:   "Invalid input (negative number)",
			n:      -1,
			result: 0,
			err:    ErrNegative,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res, err := fib(tc.n)
			if res != tc.result {
				t.Errorf("Unexpected result: got %d, want %d", res, tc.result)
			}
			if err != tc.err {
				t.Errorf("Unexpected error: got %v, want %v", err, tc.err)
			}
		})
	}
}
