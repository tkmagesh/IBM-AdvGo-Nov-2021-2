package main

import (
	"fmt"
	"time"
)

func main() {
	isPrime := memoize(isPrime)

	fmt.Println(isPrime(96))
	fmt.Println(isPrime(97))
	fmt.Println(isPrime(98))
	fmt.Println(isPrime(99))

	fmt.Println("Processing again.....")
	fmt.Println(isPrime(96))
	fmt.Println(isPrime(97))
	fmt.Println(isPrime(98))
	fmt.Println(isPrime(99))

}

func memoize(oper func(int) bool) func(int) bool {
	cache := make(map[int]bool)
	return func(n int) bool {
		if result, ok := cache[n]; ok {
			return result
		}
		result := oper(n)
		cache[n] = result
		return result
	}
}

func isPrime(n int) bool {
	fmt.Println("Processing : ", n)
	time.Sleep(5 * time.Second)
	if n <= 1 {
		return false
	}

	for i := 2; i <= (n / 2); i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
