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

	rows := strings.Split(string(content), "\n")
	checksum1 := 0
	checksum2 := 0

	for _, row := range rows {
		min := 0
		max := 0
		elements := strings.Split(row, "\t")
		for i := 0; i < len(elements); i++ {
			element, _ := strconv.Atoi(elements[i])

			// Part 2
			for j := 0; j < len(elements); j++ {
				comparator, _ := strconv.Atoi(elements[j])
				if element != comparator {
					if element % comparator == 0 {
						checksum2 += element / comparator
						continue
					}
				}
			}

			// Part 1
			if i == 0 {
				min = element
				max = element
				continue
			}

			if element > max {
				max = element
			}
			if element < min {
				min = element
			}
		}
		checksum1 += max - min
	}
	fmt.Println("Part 1:", checksum1)
	fmt.Println("Part 2:", checksum2)
}
