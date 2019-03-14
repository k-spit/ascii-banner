package main

import (
	"fmt"
	"time"
)

const input = "This is the banner text"

func main() {
	asciiBanner(input, 200)
}

func asciiBanner(input string, t time.Duration) {
	input += " - "
	i := 0
	for {
		print("\033[H\033[2J")
		input = shift(input)
		fmt.Println(input)
		time.Sleep(time.Millisecond * t)
		i++
		if i == len(input) {
			i = 0
		}
	}
}

//shift takes last character and places it as very first character
func shift(s string) string {
	return string(s[len(s)-1]) + s[:len(s)-1]
}
