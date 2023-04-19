package main

import "fmt"

func main() {
	number := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	x := 17
	result := findIndex(number, x)
	if result == -1 {
		fmt.Printf("This number is not in the array!!!")
	} else {
		fmt.Printf("Index %d is %d\n", x, result)
	}
}

func findIndex(arr []int, x int) int {

	for i := 0; i < len(arr); i++ {
		if arr[i] == x {
			return i
		}
	}
	return -1
}
