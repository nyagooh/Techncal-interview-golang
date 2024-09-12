package main

import "fmt"

func main() {
	n := 60
	factors := Factorial(n)
	fmt.Println(factors)
}

// Factorial finds all factors of a number, replacing composite factors with their prime factors
func Factorial(n int) []int {
	var factors []int
	for i := 1; i <= n; i++ {
		if n%i == 0 {
			// If the number is composite and has factors, replace it with its prime factors
			if i > 1 && HasFactors(i) {
				factors = append(factors, GetPrimeFactors(i)...)
			} else {
				factors = append(factors, i)
			}
		}
	}
	return factors
}

// GetPrimeFactors returns the prime factors of a given number
func GetPrimeFactors(n int) []int {
	var primeFactors []int
	for i := 2; i <= n; i++ {
		for n%i == 0 {
			primeFactors = append(primeFactors, i)
			n /= i
		}
	}
	return primeFactors
}

// HasFactors checks if a number has factors other than 1 and itself
func HasFactors(n int) bool {
	for i := 2; i <= n/2; i++ {
		if n%i == 0 {
			return true
		}
	}
	return false
}
