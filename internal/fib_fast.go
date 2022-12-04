//go:build fast

package internal

func fib(n int) (int, error) {
	if n < 0 {
		return 0, ErrNegative
	} else if n == 0 {
		return 0, nil
	}
	i := 1
	ans := []int{1, 1, 1, 0} // 2
	m := []int{1, 1, 1, 0}   // 2
	for i <= n {
		if (n & i) != 0 {
			mul(ans, m)
		}
		mul(m, m)
		i <<= 1
	}
	return ans[3], nil
}

// 2 x 2 * 2 x 2
func mul(b, m []int) {
	b[0], b[1], b[2], b[3] = b[0]*m[0]+b[1]*m[2],
		b[0]*m[1]+b[1]*m[3],
		b[2]*m[0]+b[3]*m[2],
		b[2]*m[1]+b[3]*m[3]
}
