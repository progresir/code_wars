package main

import "fmt"

func main() {
	fmt.Println(MaximumSubarraySum([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
}
func MaximumSubarraySum(numbers []int) int {
	resMax := 0
	sum := 0
	for j := 0; j < len(numbers); j++ {
		sum += numbers[j]
		if sum > resMax {
			resMax = sum
		}
		if sum < 0 {
			sum = 0
		}
	}
	return resMax
}
