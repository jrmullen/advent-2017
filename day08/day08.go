package main

import (
	"io/ioutil"
	"log"
	"strings"
	"fmt"
	"strconv"
)

var registers = make(map[string]int)

func main() {
	content, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	input := strings.Split(string(content), "\n")
	var amounts []string
	var conditions []string

	for _, row := range input {
		amounts = append(amounts, strings.Split(row, " if ")[0])
		conditions = append(conditions, strings.Split(row, " if ")[1])
	}

	for i, condition := range conditions {
		left := strings.Split(condition, " ")[0]
		right := strings.Split(condition, " ")[2]
		rightInt, _ := strconv.Atoi(right)
		operator := strings.Split(condition, " ")[1]

		switch operator {
		case ">":
			createEmptyRegisterIfNotExists(left)
			if registers[left] > rightInt {
				updateRegister(amounts[i])
			}
		case "<":
			createEmptyRegisterIfNotExists(left)
			if registers[left] < rightInt {
				updateRegister(amounts[i])
			}
		case "==":
			createEmptyRegisterIfNotExists(left)
			if registers[left] == rightInt {
				updateRegister(amounts[i])
			}
		case "!=":
			createEmptyRegisterIfNotExists(left)
			if registers[left] != rightInt {
				updateRegister(amounts[i])
			}
		case ">=":
			createEmptyRegisterIfNotExists(left)
			if registers[left] >= rightInt {
				updateRegister(amounts[i])
			}
		case "<=":
			createEmptyRegisterIfNotExists(left)
			if registers[left] <= rightInt {
				updateRegister(amounts[i])
			}
		}
	}

	largest := findLargest()

	fmt.Println("Part 1:", largest)
	fmt.Println("Part 2:", )

}

func updateRegister(input string) {
	element := strings.Split(input, " ")[0]
	direction := strings.Split(input, " ")[1]
	amount := strings.Split(input, " ")[2]
	amountInt, _ := strconv.Atoi(amount)

	// Check if exists in register map
	if _, exists := registers[element]; exists {
		if direction == "inc" {
			registers[element] += amountInt
		}

		if direction == "dec" {
			registers[element] -= amountInt
		}
	} else { // Create new register
		registers[element] = 0
		if direction == "inc" {
			registers[element] += amountInt
		}

		if direction == "dec" {
			registers[element] -= amountInt
		}
	}
}

func findLargest() int {
	largest := 0
	for _, register := range registers {
		if register > largest {
			largest = register
		}
	}

	return largest
}

func createEmptyRegisterIfNotExists(element string) {
	// Create new empty register if it does not already exist
	if _, exists := registers[element]; !exists {
		registers[element] = 0
	}
}
