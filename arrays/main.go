//find the largest and smallest elements in an array without using sorting functions

package main

import "fmt"

func main() {
	num := []int{1, 2, 3, 4}
	fmt.Println(FindLargestAndSmallest(num))
}

func FindLargestAndSmallest(arr []int) (int,int) {
smallest := arr[0]
laregest := arr[1]
for _,ch := range arr {
	if ch < smallest {
        smallest = ch
    }
    if ch > laregest {
        laregest = ch
    }
}
return laregest, smallest
}
