package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
	"strconv"
	"encoding/hex"
)

func main() {
	content, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	inputStrs := strings.Split(string(content), ",")
	input := make([]int, len(inputStrs))

	// Convert input to integers
	for i, element := range inputStrs {
		input[i], _ = strconv.Atoi(element)
	}

	// Populate circular list
	list := make([]int, 256)
	for i := range list {
		list[i] = i
	}

	currentPosition := 0
	skipSize := 0

	// Part 1
	for _, length := range input {

		j := currentPosition
		k := (currentPosition + length - 1) % len(list)

		// For each element in the range of length
		for i := 0; i < length/2; i++ {

			// Swap positions positions
			list[j%len(list)], list[k%len(list)] = list[k%len(list)], list[j%len(list)]

			j = (j + 1) % len(list)
			k = (k - 1 + len(list)) % len(list)
		}

		// Move the current position forward by that length plus the skip size
		currentPosition += length + skipSize

		// Increase the skip size by one
		skipSize++
	}

	fmt.Println("Part 1:", list[0]*list[1])

	// Part 2

	// Reset values
	currentPosition = 0
	skipSize = 0

	// Add the following lengths to the end of the sequence: 17, 31, 73, 47, 23
	specialSqeuence := []byte("17, 31, 73, 47, 23")
	content = append(content, specialSqeuence...) // ... allows appending a slice to another slice

	sparseHash := make([]int, 256)
	for i := range sparseHash {
		sparseHash[i] = i
	}

	for i := 0; i < 64; i++ {
		for _, length := range content {
			lengthInt := int(length)

			j := currentPosition
			k := (currentPosition + lengthInt - 1) % len(sparseHash)

			// Start at both ends and meet in the middle
			for i := 0; i < lengthInt/2; i++ {

				// Swap values
				sparseHash[j%len(sparseHash)], sparseHash[k%len(sparseHash)] = sparseHash[k%len(sparseHash)], sparseHash[j%len(sparseHash)]

				j = (j + 1) % len(sparseHash)
				k = (k - 1 + len(sparseHash)) % len(sparseHash)
			}

			// Move the current position forward by that length plus the skip size
			currentPosition += lengthInt + skipSize

			// Increase the skip size by one
			skipSize++
		}
	}

	// XOR every 16 elements
	denseHash := make([]byte,len(list)/16)
	for i:=0; i < len(list);i+=16{
		for j:=0; j <16; j++{
			denseHash[i/16] ^= byte(sparseHash[i+j])
		}
	}

	// Convert to hex
	hexString := hex.EncodeToString(denseHash)

	fmt.Println("Part 2:", hexString)
}
