//go:build lazy

package internal

func fib(n int) (int, error) {
	if n == 0 {
		return 0, nil
	} else if n == 1 {
		return 1, nil
	} else if n < 0 {
		return 0, ErrNegative
	}
	l, err := fib(n - 1)
	if err != nil {
		return 0, err
	}
	r, err := fib(n - 2)
	if err != nil {
		return 0, err
	}
	return l + r, nil
}
