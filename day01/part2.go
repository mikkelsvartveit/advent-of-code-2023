package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func replaceStringWithDigit(line string) string {
	line = strings.ReplaceAll(line, "one", "1")
	line = strings.ReplaceAll(line, "two", "2")
	line = strings.ReplaceAll(line, "three", "3")
	line = strings.ReplaceAll(line, "four", "4")
	line = strings.ReplaceAll(line, "five", "5")
	line = strings.ReplaceAll(line, "six", "6")
	line = strings.ReplaceAll(line, "seven", "7")
	line = strings.ReplaceAll(line, "eight", "8")
	line = strings.ReplaceAll(line, "nine", "9")

	return line
}

func main() {
	file, err := os.Open("day01/input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
	}

	data, err := io.ReadAll(file)
	if err != nil {
		fmt.Println("Error reading file:", err)
	}

	defer file.Close()

	lines := strings.Split(strings.TrimSpace(string(data)), "\n")

	sum := 0
	for _, line := range lines {
		var digit1, digit2 string

	loop1:
		for i := 0; i <= len(line); i++ {
			substr := line[:i]
			substr = replaceStringWithDigit(substr)

			for i := 0; i < len(substr); i++ {
				char := substr[i]
				if char >= '0' && char <= '9' {
					digit1 = string(char)
					break loop1
				}
			}
		}

	loop2:
		for i := len(line); i >= 0; i-- {
			substr := line[i:]
			substr = replaceStringWithDigit(substr)

			for i := len(substr) - 1; i >= 0; i-- {
				char := substr[i]
				if char >= '0' && char <= '9' {
					digit2 = string(char)
					break loop2
				}
			}
		}

		number, _ := strconv.Atoi(digit1 + digit2)
		sum += number
	}

	fmt.Println(sum)
}
