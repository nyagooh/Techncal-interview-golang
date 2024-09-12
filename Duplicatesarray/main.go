// remove duplicates from an  array of integers without using extra space
package main

import "fmt"

func main() {
	s := []int{2, 3, 2, 4, 5, 3, 5, 6, 2, 7, 3, 7}
	fmt.Println(removeDuplicates(s))
}

func removeDuplicates(arr []int) ([]int, int) {
	duplicates := make(map[int]bool)
	var count int
	result := []int{}
	for _, i := range arr {
		duplicates[i] = true
	}
for key := range duplicates {
		result = append(result, key)
	}
	for _, i := range arr {
		if duplicates[i] {
			count++
		}
		duplicates[i] = false
	}

	return result, count
}
