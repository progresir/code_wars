package main

import "strings"

func main() {
	println(Accum("hghj"))
}

func Accum(s string) string {
	var res string
	for i := range s {
		for j := 0; j <= i; j++ {
			if j == 0 {
				res += strings.ToUpper(string(s[i]))
			} else {
				res += strings.ToLower(string(s[i]))
			}
		}
		if i != len(s)-1 {
			res += "-"
		}
	}
	return res
}
