//go:build gen8

package internal

import (
	"errors"
	"math/big"
)

func fib(n int) (int, error) {
	if n < 0 {
		return 0, errors.New("ErrNegative")
	}

	matrix := [2][2]*big.Int{
		[2]*big.Int{big.NewInt(1), big.NewInt(1)},
		[2]*big.Int{big.NewInt(1), big.NewInt(0)},
	}

	var result *big.Int = big.NewInt(1)
	for n > 0 {
		if n&1 == 1 {
			result = result.Mul(result, matrix[0][0]).Add(result, matrix[0][1])
		}
		matrix = [2][2]*big.Int{
			[2]*big.Int{
				matrix[0][0].Mul(matrix[0][0], matrix[0][0]).Add(matrix[0][0], matrix[0][1]),
				matrix[0][1].Mul(matrix[0][0], matrix[1][0]).Add(matrix[0][1], matrix[1][1]),
			},
			[2]*big.Int{
				matrix[1][0].Mul(matrix[0][0], matrix[1][0]).Add(matrix[1][0], matrix[1][1]),
				matrix[1][1].Mul(matrix[1][0], matrix[1][0]).Add(matrix[1][1], matrix[1][1]),
			},
		}
		n >>= 1
	}
	return int(result.Int64()), nil
}
