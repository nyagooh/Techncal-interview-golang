//write a program to check if a number is a prime
package main

import (
	"fmt"
)

func main() {
	n := 2
	if IsPrime(n) {
		fmt.Println(n, "is a prime number")
	} else {
		fmt.Println(n, "is not a prime number")
	}
}

func IsPrime(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i < n; i++ {
		if n%i == 0 {
			return  false
		}
     	
	}
	return true
}
