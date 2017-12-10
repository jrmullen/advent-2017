package main

import (
	"io/ioutil"
	"log"
	"strings"
	"fmt"
	"strconv"
	"reflect"
)

func main() {
	content, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	strs := strings.Split(string(content), "\t")
	input := InputToInt(strs)

	cycles := 0
	snapshots := make([][]int, 0)
	solved := false
	var endingSnapshot []int

	// Part 1
	for { // Infinite loop
		if solved {
			fmt.Println("Part 1:", cycles)
			break
		}

		cycles++

		input = Reallocate(input)

		for _, snapshot := range snapshots {
			if reflect.DeepEqual(snapshot, input) {
				endingSnapshot = snapshot
				solved = true
				break
			}
		}

		// Add data to snapshots list
		snapshot := make([]int, len(input))
		copy(snapshot, input)
		snapshots = append(snapshots, snapshot)
	}

	cycles = 0
	solved = false

	// Part 2
	for {

		cycles++

		input = Reallocate(input)

		if reflect.DeepEqual(input, endingSnapshot) {
			fmt.Println("Part 2:", cycles)
			return
		}

	}
}

func InputToInt(strs []string) []int {
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

func Reallocate(input []int) []int {
	var indexOfLargest int
	// Find indexOfLargest bin
	for i, binVal := range input {
		if binVal > input[indexOfLargest] {
			indexOfLargest = i // Set indexOfLargest to index of indexOfLargest number
		}
	}

	highestBinValue := input[indexOfLargest]
	input[indexOfLargest] = 0 // Empty the largest bin

	for i := 1; i <= highestBinValue; i++ {
		input[(indexOfLargest+i)%len(input)]++
	}

	return input
}
