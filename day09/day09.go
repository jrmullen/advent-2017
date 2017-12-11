package main

import (
	"io/ioutil"
	"log"
	"fmt"
)

func main() {
	content, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	input := string(content)

	garbage := false
	total := 0
	depth := 0
	garbageCount := 0

	for i := 0; i < len(input); {
		char := string(input[i])
		if garbage {
			if char == ">" {
				garbage = false
				continue
			}

			if char == "!" {
				i += 2
				continue
			}

			garbageCount ++
		} else {
			if char == "<" {
				garbage = true
			}

			if char == "{" {
				depth += 1
				total += depth
			}

			if char == "}" {
				depth -= 1
			}
		}

		i++
	}

	fmt.Println("Part 1:", total)
	fmt.Println("Part 2:", garbageCount)
}
