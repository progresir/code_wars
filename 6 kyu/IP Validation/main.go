package main

import (
	"strconv"
	"strings"
)

func main() {
	//println(Is_valid_ip("1223.045.067.089"))
	println(Is_valid_ip("123.045.067.089"))
}

func Is_valid_ip(ip string) bool {
	mass := strings.Split(ip, ".")
	if len(mass) != 4 {
		return false
	}
	for _, c := range mass {
		if len(c) > 3 && (len(c) > 1 && string(c[0]) == "0") {
			return false
		}
		if len(c) > 1 && string(c[0]) == "0" {
			return false
		}
		cInt, err := strconv.Atoi(c)
		if err != nil {
			return false
		}
		if !(cInt >= 0 && cInt <= 255) {
			return false
		}
	}
	return true
}
