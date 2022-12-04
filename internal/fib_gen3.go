//go:build gen3

package internal

func fib(n int) (int, error) {
	if n < 0 {
		return 0, ErrNegative
	} else if n == 0 {
		return 0, nil
	} else if n == 1 {
		return 1, nil
	}
	n1, _ := fib(n - 1)
	n2, _ := fib(n - 2)
	return n1 + n2, nil
}
