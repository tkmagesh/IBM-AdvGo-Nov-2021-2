package main

func main() {

}

func IsDivisible(n1, n2 int) bool {
	return n1%n2 == 0
}

func IsPrime(n int) bool {
	if n <= 1 {
		return false
	}
	if n <= 3 {
		return true
	}
	for idx := 2; idx < n; idx++ {
		if IsDivisible(n, idx) {
			return false
		}
	}
	return true
}

func GeneratePrimes(start, end int) []int {
	var primes []int = make([]int, 100)
	var count int
	for idx := start; idx <= end; idx++ {
		if IsPrime(idx) {
			primes[count] = idx
			count++
		}
	}
	return primes[:count]
}
