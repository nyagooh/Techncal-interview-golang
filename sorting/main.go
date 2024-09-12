//challenge write a sorting algorithim

package main

import "fmt"

func main() {
	num := []int{2, 4, 1, 3, 5, 6, 7, 3, 9}
	fmt.Println(Sort(num))
}
func Sort(s []int) []int {
	len := len(s)
	for i := 0; i < len-1; i++ {
		for j := 0; j < len-1-i; j++ {
			if s[j] > s[j+1] {
				s[j], s[j+1] = s[j+1], s[j]
			}
		}
	}
	return s
}
