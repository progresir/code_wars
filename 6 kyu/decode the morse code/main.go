package main

import "strings"

var MORSE_CODE = map[string]string{
	".-":        "A",
	"-...":      "B",
	"-.-.":      "C",
	"-..":       "D",
	".":         "E",
	"..-.":      "F",
	"--.":       "G",
	"....":      "H",
	"..":        "I",
	".---":      "J",
	"-.-":       "K",
	".-..":      "L",
	"--":        "M",
	"-.":        "N",
	"---":       "O",
	".--.":      "P",
	"--.-":      "Q",
	".-.":       "R",
	"...":       "S",
	"-":         "T",
	"..-":       "U",
	"...-":      "V",
	".--":       "W",
	"-..-":      "X",
	"-.--":      "Y",
	"--..":      "Z",
	"...---...": "SOS",
	"-----":     "0",
	".----":     "1",
	"..---":     "2",
	"...--":     "3",
	"....-":     "4",
	".....":     "5",
	"-....":     "6",
	"--...":     "7",
	"---..":     "8",
	"----.":     "9",
	".-.-.-":    ".",
	"--..--":    ",",
	"..--..":    "?",
	"-.-.--":    "!",
	"---...":    ":",
	".-..-.":    "\"",
	".----.":    "'",
	"-...-":     "=",
}

func main() {
	println("|" + DecodeMorse("      ...---... -.-.--   - .... .   --.- ..- .. -.-. -.-   -... .-. --- .-- -.   ..-. --- -..-   .--- ..- -- .--. ...   --- ...- . .-.   - .... .   .-.. .- --.. -.--   -.. --- --. .-.-.-  ") + "|")
}
func DecodeMorse(morseCode string) (res string) {
	massWords := strings.Split(morseCode, "   ")
	for i := range massWords {
		if i != 0 && res != "" {
			res += " "
		}
		massLetters := strings.Split(massWords[i], " ")
		for j := range massLetters {

			if val, ok := MORSE_CODE[massLetters[j]]; ok {
				res += val
			}
		}

	}
	return strings.TrimSpace(res)
}
