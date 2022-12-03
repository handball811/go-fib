//go:build normal

package internal

func fib(n int) (int, error) {
	if n == 0 {
		return 0, nil
	} else if n < 0 {
		return 0, ErrNegative
	}
	a := 0
	b := 1
	for i := 0; i < n-1; i++ {
		b += a
		a = b - a
	}
	return b, nil
}
