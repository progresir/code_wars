package main

import "fmt"

func Josephus(items []interface{}, k int) []interface{} {
	res := make([]interface{}, 0)
	index := 0
	count := 0
	for len(items) != count {
		for i := k - 1; i > -1; {
			if index >= len(items) {
				index = 0
			}
			if items[index] == -1 {
				index++
				continue
			}
			if i == 0 {
				res = append(res, items[index])
				items[index] = -1
				count++
			}
			index++
			i--
		}
	}
	return res
}

func main() {

	fmt.Println(Josephus([]interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 2))
}
