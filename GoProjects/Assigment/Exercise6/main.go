package main

import (
	"fmt"
	"reflect"
)

func main() {
	nums := []int{1, 2, 3, 2, 1}
	swap := reflect.Swapper(nums)

	//swapping
	fmt.Println("nums array before swapping ", nums)
	swap(0, 1)
	swap(2, 3)
	fmt.Println("nums array after swapping ", nums)

	//reversing nums
	fmt.Println("nums array before reversing ", nums)
	for i, j := 0, len(nums)-1; i < j; i, j = i+1, j-1 {
		swap(i, j)
	}
	fmt.Println("nums array after reversing ", nums)

	//comparing swapped array to new array
	newNums := []int{1, 3, 2, 1, 2}
	isEqual := reflect.DeepEqual(nums, newNums)

	if isEqual {
		fmt.Println("Both slices are equal")
	} else {
		fmt.Println("Both slices are not equal")
	}
}
