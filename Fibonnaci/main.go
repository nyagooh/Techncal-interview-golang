//write a fibonnaci sequence

package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(Fibonacci(10))
}
func Fibonacci(n int) []int {
	fibSeries := make([]int, n)
	if n > 1 {
		fibSeries[0] = 0
		fibSeries[1] = 1
		for i := 2; i < n; i++ {
			fibSeries[i] = fibSeries[i-1] + fibSeries[i-2]
		}
	} else {
		os.Exit(0)
	}
	return fibSeries
}
