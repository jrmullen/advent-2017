package main

import (
	"io/ioutil"
	"log"
	"strings"
	"fmt"
	"strconv"
)

func main() {
	content, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	strs := strings.Split(string(content), "\n")

	input := SetInput(strs)
	steps1 := Part1(input)

	input = SetInput(strs) //TODO is input being passed by reference to Part1()?
	steps2 := Part2(input)

	fmt.Println("Part 1:", steps1)
	fmt.Println("Part 2:", steps2)

}

func Part1(input []int) int {
	steps := 0
	current := 0
	for i := 0; i < len(input); {
		steps++
		current += input[i]
		input[i] = input[i] + 1
		i = current

		if i > len(input) {
			break
		}
	}

	return steps
}

func Part2(input []int) int {
	steps := 0
	current := 0
	for i := 0; i < len(input); {
		steps++
		current += input[i]

		if input[i] >= 3 {
			input[i] = input[i] - 1
		} else {
			input[i] = input[i] + 1
		}
		i = current

		if i > len(input) {
			break
		}
	}

	return steps
}

func SetInput(strs []string) []int {
	input := []int{}
	// Convert strings to int
	for _, element := range strs {
		k, err := strconv.Atoi(element)
		if err != nil {
			panic(err)
		}
		input = append(input, k)
	}

	return input
}
