package main

import (
	"io/ioutil"
	"log"
	"strings"
	"fmt"
)

func main() {
	content, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	input := strings.Split(string(content), ",")

	fmt.Println(input)
}