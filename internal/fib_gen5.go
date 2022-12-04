//go:build gen5

package internal

// fib is a faster version of the fib function that uses memoization
// to store previously calculated values and avoid unnecessary repeated calculations.
func fib(n int) (int, error) {
	if n < 0 {
		return 0, ErrNegative
	}

	memo := map[int]int{}

	var fibRecursive func(int) int
	fibRecursive = func(n int) int {
		if v, ok := memo[n]; ok {
			return v
		}
		if n == 0 {
			return 0
		} else if n == 1 {
			return 1
		}

		v := fibRecursive(n-1) + fibRecursive(n-2)
		memo[n] = v
		return v
	}

	return fibRecursive(n), nil
}
