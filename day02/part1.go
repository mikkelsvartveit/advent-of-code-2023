package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("day02/input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
	}

	data, err := io.ReadAll(file)
	if err != nil {
		fmt.Println("Error reading file:", err)
	}

	defer file.Close()

	lines := strings.Split(strings.TrimSpace(string(data)), "\n")

	idSum := 0
gameLoop:
	for i, line := range lines {
		lineStripped := strings.TrimSpace(strings.Split(line, ": ")[1])

		gameRounds := strings.Split(lineStripped, "; ")

		for _, round := range gameRounds {
			colorsAndNumbers := strings.Split(strings.TrimSpace(round), ", ")

			for _, colorAndNumberString := range colorsAndNumbers {
				colorAndNumberSlice := strings.Split(colorAndNumberString, " ")
				number, _ := strconv.Atoi(colorAndNumberSlice[0])
				color := colorAndNumberSlice[1]

				if (color == "red" && number > 12) || (color == "green" && number > 13) || (color == "blue" && number > 14) {
					continue gameLoop
				}
			}
		}

		idSum += i + 1
	}

	println(idSum)
}
