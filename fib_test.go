package fib

import "testing"

func BenchmarkFibonacci_10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Fibonacci(10)
	}
}

func BenchmarkFibonacci_20(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Fibonacci(20)
	}
}

func BenchmarkFibonacciBig_10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FibonacciBig(10)
	}
}

func BenchmarkFibonacciBig_20(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FibonacciBig(20)
	}
}

func TestFibonacci(t *testing.T) {
	data := []struct {
		n    uint
		want uint64
	}{
		{0, 0}, {1, 1}, {2, 1}, {3, 2}, {4, 3}, {5, 5}, {6, 8}, {10, 55}, {42, 267914296},
	}

	for _, d := range data {
		if got := Fibonacci(d.n); got != d.want {
			t.Errorf("Invalid Fibonacci value for N: %d, got: %d, want: %d", d.n, got, d.want)
		}
	}
}

func TestFibonacciBig(t *testing.T) {
	data := []struct {
		n    uint
		want int64
	}{
		{0, 0}, {1, 1}, {2, 1}, {3, 2}, {4, 3}, {5, 5}, {6, 8}, {10, 55}, {42, 267914296},
	}

	for _, d := range data {
		if got := FibonacciBig(d.n); got.Int64() != d.want {
			t.Errorf("Invalid Fibonacci value for N: %d, got: %d, want: %d", d.n, got, d.want)
		}
	}
}
