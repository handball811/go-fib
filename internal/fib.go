package internal

import "errors"

var (
	ErrNegative = errors.New("negative number")
)

func Run(n int) (int, error) {
	return fib(n)
}
