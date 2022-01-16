package main

import "fmt"

func main() {
	d := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	//d2 := [][]int{{}}
	fmt.Println(Snail(d))
}

func Snail(snaipMap [][]int) []int {
	res := make([]int, 0)
	glob := 0
	lens := len(snaipMap)
	for {
		for i := 0 + glob; lens > i+glob; i++ {
			if len(snaipMap[glob]) == 0 {
				return res
			}
			res = append(res, snaipMap[glob][i])
		}
		for i := 1 + glob; lens > i+glob; i++ {
			res = append(res, snaipMap[i][lens-1-glob])
		}
		for i := lens - 1 - glob - 1; i >= glob; i-- {
			res = append(res, snaipMap[lens-1-glob][i])
		}
		for i := lens - 1 - glob - 1; i > glob; i-- {
			res = append(res, snaipMap[i][glob])
		}
		glob++
		if glob >= lens {
			break
		}
	}
	return res
}
