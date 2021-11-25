package main

func Multiple3And5(number int) int {
	res := 0
	if number < 0 {
		return 0
	}
	for i := 1; i < number; i++ {
		if i%3 == 0 {
			res += i
			continue
		}
		if i%5 == 0 {
			res += i
		}
		println(res)
	}
	return res
}

func main() {
	println(Multiple3And5(10))
}
