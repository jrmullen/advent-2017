package main

import (
	"io/ioutil"
	"log"
	"fmt"
	"strconv"
	"math"
)

type Coord struct {
	x, y int
}

var spiral = make(map[Coord]int)

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
	var posX int
	var posY int
	var largest int
	var part2 int
	part2Answered := false

	for steps < n-1 {
		length := (turns / 2) + 1
		for i := 0; i < length; i++ {
			if steps == n-1 {
				break
			}

			if i == 0 {
				spiral[Coord{0, 0}] = 1
			}

			largest = CalculateValue(posX, posY)
			spiral[Coord{posX, posY}] = largest

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

			if largest > n && !part2Answered {
				part2Answered = true
				part2 = largest
			}

		}

		turns++
	}

	fmt.Println("Part 1:", math.Abs(float64(posX))+math.Abs(float64(posY)))
	fmt.Println("Part 2:", part2)
}

func CalculateValue(x, y int) int {
	var total int

	total = spiral[Coord{x, y}] +
		spiral[Coord{x + 1, y}] +
		spiral[Coord{x + 1, y + 1}] +
		spiral[Coord{x, y + 1}] +
		spiral[Coord{x - 1, y + 1}] +
		spiral[Coord{x - 1, y}] +
		spiral[Coord{x - 1, y - 1}] +
		spiral[Coord{x, y - 1}] +
		spiral[Coord{x + 1, y - 1}]

	return total
}
