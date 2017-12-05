package main

import (
	"io/ioutil"
	"log"
	"fmt"
	"strings"
	"sort"
)

func main() {
	content, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	input := strings.Split(string(content), "\n")
	validCount1 := 0
	validCount2 := 0

	for _, row := range input {
		words := strings.Split(row, " ")
		set := make(map[string]bool)
		valid := true
		anagram := ContainsAnagram(row) // Part 2
		for _, word := range words {
			if !set[word] {
				set[word] = true
			} else {
				valid = false
				continue
			}
		}

		if valid {
			validCount1++
		}

		if !anagram {
			validCount2++
		}
	}

	fmt.Println("Part 1:", validCount1)
	fmt.Println("Part 2:", validCount2)
}

func ContainsAnagram(row string) bool {
	words := strings.Split(row, " ")
	for x := 0; x < len(words); x++ {

		word := StringToRunes(words[x])

		sort.Slice(word, func(i, j int) bool { // Sort the slice of first word's runes
			return word[i] < word[j]
		})

		for y := 0; y < len(words); y++ {
			if y == x {
				continue // Skip the word it is the word currently being used
			}
			comparator := StringToRunes(words[y])

			sort.Slice(comparator, func(i, j int) bool {
				return comparator[i] < comparator[j]
			})

			if string(word) == string(comparator) {
				return true
			}
		}
	}

	return false
}

func StringToRunes(s string) []int32 {
	var runes []int32
	for _, r := range s {
		runes = append(runes, r)
	}

	return runes
}
