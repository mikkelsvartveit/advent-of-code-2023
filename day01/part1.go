package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

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

		for i := 0; i < len(line); i++ {
			char := line[i]
			if char >= '0' && char <= '9' {
				digit1 = string(char)
				break
			}
		}

		for i := len(line) - 1; i >= 0; i-- {
			char := line[i]
			if char >= '0' && char <= '9' {
				digit2 = string(char)
				break
			}
		}

		number, _ := strconv.Atoi(digit1 + digit2)
		sum += number
	}

	fmt.Println(sum)
}
