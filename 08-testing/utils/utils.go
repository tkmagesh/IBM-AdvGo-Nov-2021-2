package utils

import "math"

func IsPrime(n int) bool {
	if n <= 1 {
		return false
	}
	if n <= 3 {
		return true
	}
	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func IsPrime2(n int) bool {
	if n <= 1 {
		return false
	}
	if n <= 3 {
		return true
	}
	for i := 2; i <= (n / 2); i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func IsPrime3(n int) bool {
	if n <= 1 {
		return false
	}
	if n <= 3 {
		return true
	}
	var i float64

	for i = 2; i <= math.Sqrt(float64(n)); i++ {
		if n%int(i) == 0 {
			return false
		}
	}
	return true
}
