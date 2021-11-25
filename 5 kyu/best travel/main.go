package main

import (
	"fmt"
)

func main() {
	var ts = []int{91, 74, 73, 85, 73, 81, 87}
	fmt.Println(abc(331, 5, ts))

}

func abc(t, k int, ls []int) int {
	if k == 0 {
		return 0
	}
	max := -1
	for i := 0; i <= len(ls)-k; i++ {
		number := ls[len(ls)-1-i]
		j := abc(t-number, k-1, ls[:len(ls)-1-i])
		if j == -1 {
			continue
		}
		j = j + number
		println(j, " ", ls[len(ls)-1-i])
		if j <= t {
			if j > max {
				max = j
			}
		} else {
			continue
		}
	}
	return max
}
