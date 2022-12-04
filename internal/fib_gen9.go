//go:build gen9

package internal

func fib(n int) (int, error) {
	if n < 0 {
		return 0, ErrNegative
	}

	matrix := [2][2]int{{1, 1}, {1, 0}}
	if n == 0 {
		return 0, nil
	} else if n == 1 {
		return 1, nil
	}

	result := [2][2]int{{1, 1}, {1, 0}}
	n -= 2

	for n > 0 {
		if n%2 == 1 {
			result = multiply(result, matrix)
		}
		matrix = multiply(matrix, matrix)
		n /= 2
	}

	return result[0][0], nil
}

func multiply(m1, m2 [2][2]int) [2][2]int {
	return [2][2]int{
		{m1[0][0]*m2[0][0] + m1[0][1]*m2[1][0], m1[0][0]*m2[0][1] + m1[0][1]*m2[1][1]},
		{m1[1][0]*m2[0][0] + m1[1][1]*m2[1][0], m1[1][0]*m2[0][1] + m1[1][1]*m2[1][1]},
	}
}
