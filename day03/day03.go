package main

import (
	"io/ioutil"
	"log"
	"fmt"
	"strconv"
	"math"
)

func main() {
	content, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	n, err := strconv.Atoi(string(content))
	if err != nil {
		log.Fatal(err)
	}

	var steps int
	var turns int
	var posX float64
	var posY float64

	for steps < n-1 {
		length := (turns / 2) + 1
		for i := 0; i < length; i++ {
			if steps == n-1 {
				break
			}
			steps++
			direction := turns % 4
			switch direction {
			case 0:
				posX++
			case 1:
				posY++
			case 2:
				posX--
			default:
				posY--
			}
		}

		turns++
	}

	fmt.Println("Part 1:", math.Abs(posX)+math.Abs(posY))
}
