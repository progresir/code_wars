package main

func main() {
	println(Evaporator(10, 10, 5))
}

func Evaporator(content float64, evapPerDay int, threshold int) int {
	var n int
	k := content * float64(threshold) / float64(100)
	for n = 0; k <= content; n++ {
		content = content * float64(100-evapPerDay) / float64(100)
	}
	return n
}
