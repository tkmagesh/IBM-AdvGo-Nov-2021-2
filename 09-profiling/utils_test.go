package main

import "testing"

func BenchmarkGeneratePrimes_100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GeneratePrimes(3, 100)
	}
}
