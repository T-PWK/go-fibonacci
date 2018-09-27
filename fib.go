package fib

import "math/big"

// Fibonacci calculates Fibonacci number.
// This function generated correct values from 0 to 93 sequence number.
// For bigger values use FibonacciBig function
func Fibonacci(n uint) uint64 {
	if n <= 1 {
		return uint64(n)
	}

	var n2, n1 uint64 = 0, 1

	for i := uint(2); i < n; i++ {
		n2, n1 = n1, n1+n2
	}

	return n2 + n1
}

// FibonacciBig calculates Fibonacci number using bit.Int.
// For the sequence numbers below 94, it is recommended to use Fibonacci function.
func FibonacciBig(n uint) *big.Int {
	if n <= 1 {
		return big.NewInt(int64(n))
	}

	var n2, n1, tmp = big.NewInt(0), big.NewInt(1), big.NewInt(0)

	for i := uint(1); i < n; i++ {
		tmp.Set(n1)
		n1.Add(n1, n2)
		n2.Set(tmp)
	}

	return n1
}