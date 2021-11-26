package main

import (
	"strconv"
)

var arrRemainders = []int{1, 10, 9, 12, 3, 4}

func Thirt(n int) int {
	res := -1
	count := 0
	for res != count {
		res = count
		count = 0
		nStr := strconv.Itoa(n)
		for i := range nStr {
			count = count + arrRemainders[i%6]*int(nStr[len(nStr)-i-1]-'0')
		}
		n = count
	}
	return res
}
