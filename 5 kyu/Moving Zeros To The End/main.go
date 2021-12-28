package main

import "fmt"

func main() {
	fmt.Println(MoveZeros([]int{9, 0, 0, 9, 1, 2, 0, 1, 0, 1, 0, 3, 0, 1, 9, 0, 0, 0, 0, 9}))
}

func MoveZeros(arr []int) []int {
	res := make([]int, len(arr))
	ind := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] != 0 {
			res[ind] = arr[i]
			ind++
		}
	}
	return res
}
