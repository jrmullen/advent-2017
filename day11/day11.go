package main

import (
	"io/ioutil"
	"log"
	"strings"
	"fmt"
	"math"
)

func main() {
	content, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	input := strings.Split(string(content), ",")
	x := 0
	y := 0
	var distance float64
	var furthest float64

	// Reference for axial coordinates https://www.redblobgames.com/grids/hexagons/
	// Under the Axial Coordinates section
	for _, direction := range input {
		switch direction {
		case "n":
			y--

		case "ne":
			x++
			y--

		case "se":
			x++

		case "s":
			y++

		case "sw":
			x--
			y++

		case "nw":
			x--
		}

		distance = (math.Abs(float64(x)) + math.Abs(float64(y)) + math.Abs(float64(x)+float64(y))) / 2

		if distance > furthest {
			furthest = distance
		}
	}

	fmt.Println("Part 1:", distance)
	fmt.Println("Part 2:", furthest)
}
