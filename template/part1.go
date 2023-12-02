package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("day00/input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
	}

	data, err := io.ReadAll(file)
	if err != nil {
		fmt.Println("Error reading file:", err)
	}

	defer file.Close()

	lines := strings.Split(strings.TrimSpace(string(data)), "\n")

	fmt.Println(len(lines))
}
