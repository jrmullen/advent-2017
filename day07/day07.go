package main

import (
	"io/ioutil"
	"log"
	"strings"
	"strconv"
	"fmt"
)

type node struct {
	name     string
	weight   int
	parent   string
	children []string
}

var nodes = make(map[string]node)

func main() {
	content, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	input := strings.Split(string(content), "\n")

	// Populate nodes
	for _, line := range input {
		populateNodeData(line)
	}

	// The root node's name
	var root string

	// Part 1
	for _, node := range nodes {
		if node.parent == "" {
			root = node.name
			fmt.Println("Part 1:", root)
			break
		}
	}

	// Part 2
	correctedWeight := findCorrectedWeight(root)
	fmt.Println("Part 2:", correctedWeight)

}

func populateNodeData(line string) {
	// Filter data
	name := strings.Split(line, " ")[0]
	weight := strings.Split(line, " ")[1]
	weightInt, _ := strconv.Atoi(weight[1: len(weight)-1])
	var children []string

	if strings.Contains(line, "->") {
		children = strings.Split(line, "-> ")
		children = strings.Split(children[1], ", ")
	}

	var currentNode node

	// Check if already exists
	if v, exists := nodes[name]; exists {
		currentNode = v
	}

	currentNode.name = name
	currentNode.weight = weightInt

	// Populate children
	for _, child := range children {
		if v, exists := nodes[child]; exists {
			// Update child node
			v.parent = currentNode.name
			nodes[child] = v
		} else {
			// Create new child node
			newChildNode := node{
				name:   child,
				parent: currentNode.name, // No parent
			}

			nodes[child] = newChildNode
		}

		// Add the child
		currentNode.children = append(currentNode.children, child)
	}

	// Add the new node to the list
	nodes[name] = currentNode
}

func calculateWeight(input map[string]node, root string) int {
	var sum = 0

	// Sum the weights of every child
	for _, child := range input[root].children {
		sum += calculateWeight(input, child)
	}

	return sum + input[root].weight
}

func findCorrectedWeight(root string) int {
	// Not the root
	if incorrectNodeName, difference := findDifference(nodes, root); incorrectNodeName != "" {

		// Found the deepest child
		if findCorrectedWeight(incorrectNodeName) == 0 {

			// Correct the weight to what it should be
			return nodes[incorrectNodeName].weight + difference
		}

		// Narrow the search and try again
		return findCorrectedWeight(incorrectNodeName)
	}

	return 0
}

func findDifference(element map[string]node, root string) (string, int) {
	// Track what has been found
	nodeSums := make(map[int]string)
	seen := make(map[int]int)

	// Calculate the weight for each child
	for _, element := range element[root].children {
		sum := calculateWeight(nodes, element)
		nodeSums[sum] = element
		seen[sum]++
	}

	var oddball string
	var normal int
	var incorrectValue int
	for value, count := range seen {

		// Incorrect value only appears once
		if count == 1 {
			oddball = nodeSums[value]
			incorrectValue = value
		} else {
			normal = value
		}
	}

	// Find the difference
	difference := normal - incorrectValue

	return oddball, difference
}
