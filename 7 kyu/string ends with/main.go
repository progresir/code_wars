package main

func solution(str, ending string) bool {
	if len(str) < len(ending) {
		return false
	}
	if str[(len(str)-len(ending)):] == ending {
		return true
	}
	return false
}

func main() {
	println(solution("abc", "bc"))
}
