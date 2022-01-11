package main

import "strings"

func main() {
	println(FirstNonRepeating("sTreSS"))
}

func FirstNonRepeating(str string) string {
	//charCopy := make(map[string]bool)
	strLower := strings.ToLower(str)
	for _, val := range str {
		if strings.Count(strLower, strings.ToLower(string(val))) < 2 {
			return string(val)
		}
	}
	return ""
}
