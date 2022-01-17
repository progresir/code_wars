package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(Solequa2(9000000041))
	//println(450000021*450000021 - 4*225000010*225000010)
}

func Solequa1(n int) [][]int {
	println(n)
	res := make([][]int, 0)
	for y := 0; y < n; y++ {
		sum := n + 4*y*y
		x := math.Sqrt(float64(sum))
		if x == float64(int(x)) {
			if n == int(x)*int(x)-4*y*y {
				//res = append(res, []int{int(x), y})
				if len(res) == 0 {
					res = append(res, []int{int(x), y})
				} else {
					res = append(res[:1], res...)
					res[0] = []int{int(x), y}
				}
			}
		}
	}
	return res
}

func Solequa2(n int) [][]int {
	println(n)
	res := make([][]int, 0)
	for i := 1; i < int(math.Sqrt(float64(n)))+1; i++ {
		if n%i != 0 {
			continue
		}
		j := n / i
		if (i+j)%2 == 0 && (j-i)%4 == 0 {
			x := (i + j) / 2
			y := (j - i) / 4
			res = append(res, []int{int(x), y})

		}
	}
	return res
}

func Solequa(n int) [][]int {
	println(n)
	res := make([][]int, 0)
	xLast := 0
	nSqrt := int(math.Sqrt(float64(n)))
	for y := xLast; ; y++ {
		for x := 0; ; x++ {
			if y == 0 && x == 0 {
				continue
			}
			if x == 0 {
				x = 2 * y
				continue
			}
			y2 := 2 * y
			if x-y2 >= nSqrt {
				xLast = x
				break
			}
			xNAy := (x - y2) * (x + y2)
			//nDel := n / xNAy
			//if nDel > 2 {
			//	x = x * 2
			//}
			if n/xNAy < 1 {
				xLast = x
				break
			}
			if n != xNAy {
				continue
			}
			if len(res) == 0 {
				res = append(res, []int{x, y})
			} else {
				res = append(res[:1], res...)
				res[0] = []int{x, y}
			}
			y = x
		}
		if y != 0 && n/(xLast+2*y) < 1 {
			break
		}
	}
	return res
}

func getabs(x float64) float64 {
	if x < 0 {
		return -x
	}

	if x == 0 {
		return 0
	}

	return x
}

func Sqrt(x float64) float64 {
	z := 1.0

	if x < 0 {
		return 0
	} else if x == 0 {
		return 0
	} else {
		for getabs(z*z-x) > 1e-6 {
			z = (z + x/z) / 2
		}
		return z
	}
}
