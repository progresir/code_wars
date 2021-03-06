package main

import (
	"fmt"
	"sort"
)

func Josephus(items []interface{}, k int) []interface{} {
	res := make([]interface{}, 0)
	index := 0
	for len(items) > 0 {
		for i := k; i >= 0; {
			if index >= len(items) {
				index = 0
			}
			if i == 0 {
				res = append(res, items[index])
				items = append(items[:index], items[index+1:]...)
				index--
			}
			index++
			i--
		}
	}
	return res
}

func main() {
	//println(getResult([]int{0, 0, 0}, []string{"-", "//", "+"}))
	println(getResult2([]int{3, 2, 4, 5}, 7))

	fmt.Println(Josephus([]interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 2))
}

func getResult2(arr []int, w int) bool {
	sort.Ints(arr)
	operMax := arr[len(arr)-1]
	countOper := 0
	res := 0
	isDivOper := false
	for {
		sort.Ints(arr)

		countOper++

		if len(arr)-1 >= 0 {
			if !isDivOper {
				if arr[len(arr)-1-1] > operMax/2 {
					operMax = arr[len(arr)-1-1]
				} else {
					operMax = operMax / 2
					isDivOper = true
				}
			}
		}

		arr[len(arr)-1] = arr[len(arr)-1] / 2
		if countOper > operMax {
			break
		}
	}
	for i := range arr {
		res = res + arr[i]
	}
	if res <= w {
		return true
	}
	return false
}

//можно оптимизировать но это время)
func getResult(numb []int, arith []string) int {
	if len(numb) == 0 {
		return 0
	}
	isResMax := false
	res := 0
	resMax := 0
	for i := 0; len(arith) > i; i++ {
		res = 0
		switch arith[i] {
		case "+":
			res = res + numb[0]
			arithCopy := make([]string, len(arith))
			copy(arithCopy, arith)
			arithCopy = append(arithCopy[:i], arithCopy[i+1:]...)
			numbCopy := make([]int, len(numb))
			copy(numbCopy, numb)
			recRec, ok := recur(numbCopy, arithCopy, res)
			if ok {
				res = recRec
			}
			if !isResMax {
				isResMax = true
				resMax = res
			} else {
				if res > resMax {
					resMax = res
				}
			}

		case "-":
			res = res - numb[0]
			arithCopy := make([]string, len(arith))
			copy(arithCopy, arith)
			arithCopy = append(arithCopy[:i], arithCopy[i+1:]...)
			numbCopy := make([]int, len(numb))
			copy(numbCopy, numb)
			recRec, ok := recur(numbCopy, arithCopy, res)
			if ok {
				res = recRec
			}
			if !isResMax {
				isResMax = true
				resMax = res
			} else {
				if res > resMax {
					resMax = res
				}
			}
		case "//":
			if res%numb[0] != 0 {
				continue
			}
			if numb[0] == 0 {
				continue
			}
			res = res / numb[0]
			arithCopy := make([]string, len(arith))
			copy(arithCopy, arith)
			arithCopy = append(arithCopy[:i], arithCopy[i+1:]...)
			numbCopy := make([]int, len(numb))
			copy(numbCopy, numb)
			recRec, ok := recur(numbCopy, arithCopy, res)
			if ok {
				res = recRec
			}
			if !isResMax {
				isResMax = true
				resMax = res
			} else {
				if res > resMax {
					resMax = res
				}
			}
		case "*":
			arithCopy := make([]string, len(arith))
			copy(arithCopy, arith)
			arithCopy = append(arithCopy[:i], arithCopy[i+1:]...)
			numbCopy := make([]int, len(numb))
			copy(numbCopy, numb)
			recRec, ok := recur(numbCopy, arithCopy, res)
			if ok {
				res = recRec
			}

			if !isResMax {
				isResMax = true
				resMax = res
			} else {
				if res > resMax {
					resMax = res
				}
			}
		}
	}
	return resMax
}

func recur(numb []int, arith []string, res int) (int, bool) {
	numb = append(numb[:0], numb[1:]...)
	if len(numb) == 0 {
		return 0, false
	}

	isResMax := false
	resMax := 0
	for i := 0; len(arith) > i; i++ {
		switch arith[i] {
		case "+":
			resR := res + numb[0]
			arithCopy := make([]string, len(arith))
			copy(arithCopy, arith)
			arithCopy = append(arithCopy[:i], arithCopy[i+1:]...)
			numbCopy := make([]int, len(numb))
			copy(numbCopy, numb)
			recRec, ok := recur(numbCopy, arithCopy, resR)
			if ok {
				resR = recRec
			}

			if !isResMax {
				isResMax = true
				resMax = resR
			} else {
				if resR > resMax {
					resMax = resR
				}
			}
		case "-":
			resR := res - numb[0]
			arithCopy := make([]string, len(arith))
			copy(arithCopy, arith)
			arithCopy = append(arithCopy[:i], arithCopy[i+1:]...)
			numbCopy := make([]int, len(numb))
			copy(numbCopy, numb)
			recRec, ok := recur(numbCopy, arithCopy, resR)
			if ok {
				resR = recRec
			}

			if !isResMax {
				isResMax = true
				resMax = resR
			} else {
				if resR > resMax {
					resMax = resR
				}
			}

		case "//":
			if res%numb[0] != 0 {
				continue
			}
			if numb[0] == 0 {
				continue
			}
			resR := res / numb[0]
			arithCopy := make([]string, len(arith))
			copy(arithCopy, arith)
			arithCopy = append(arithCopy[:i], arithCopy[i+1:]...)
			numbCopy := make([]int, len(numb))
			copy(numbCopy, numb)
			recRec, ok := recur(numbCopy, arithCopy, resR)
			if ok {
				resR = recRec
			}

			if !isResMax {
				isResMax = true
				resMax = resR
			} else {
				if resR > resMax {
					resMax = resR
				}
			}
		case "*":
			resR := res * numb[0]
			arithCopy := make([]string, len(arith))
			copy(arithCopy, arith)
			arithCopy = append(arithCopy[:i], arithCopy[i+1:]...)
			numbCopy := make([]int, len(numb))
			copy(numbCopy, numb)
			recRec, ok := recur(numbCopy, arithCopy, resR)
			if ok {
				resR = recRec
			}

			if !isResMax {
				isResMax = true
				resMax = resR
			} else {
				if resR > resMax {
					resMax = resR
				}
			}
		}
	}
	return resMax, isResMax
}
