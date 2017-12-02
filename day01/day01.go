package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"log"
)

func main() {
	content, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	var input = string(content)

	part1Sum := 0
	part2Sum := 0
	offset := len(input) / 2

	// Part 1
	if input[len(input)-1] == input[0] {
		intVal, _ := strconv.Atoi(string(input[len(input)-1]))
		part1Sum += intVal
	}

	for i := 0; i < len(input)-1; {
		// Part 1
		if input[i] == input[i+1] {
			intVal, _ := strconv.Atoi(string(input[i]))
			part1Sum += intVal

		}

		// Part 2
		comparator := i + offset
		if comparator >= len(input) {
			comparator -= len(input)
		}

		if input[i] == input[comparator] {
			intVal, _ := strconv.Atoi(string(input[i]))
			part2Sum += intVal
		}
		i++
	}

	fmt.Println("Part 1:", part1Sum)
	fmt.Println("Part 2:", part2Sum)

}
