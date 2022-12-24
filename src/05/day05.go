package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func parseFile(filename string) []string {

	data, _ := os.ReadFile(filename)
	stringData := string(data)
	return strings.Split(stringData, "\r\n")
}

type stack struct {
	elements []string
	top      int
} // TODO : we might need to have the default value as -1

func main() {

	data := parseFile("../../input/test/day05.txt")

	var stackData []string
	var moveData []string

	var split bool = false

	for i := 0; i < len(data); i++ {

		if data[i] == "" {
			split = true
		}

		if !split {
			stackData = append(stackData, data[i])
		} else if data[i] != "" {
			moveData = append(moveData, data[i])

		}

	}

	results := strings.Split(strings.TrimSpace(stackData[len(stackData)-1]), "   ")
	numberOfStacks, _ := strconv.Atoi(results[len(results)-1])

	fmt.Printf("We want to create %v stacks \n", numberOfStacks)
	// Create a list of lists of 9 elements

	var stacks = make([]stack, numberOfStacks)

	var step int = 4

	for i := len(stackData) - 2; i >= 0; i-- {
		for j := 1; j < len(stackData[i]); j += step {
			character := stackData[i][j]

			if character != 32 { // whitespace
				fmt.Printf("Trying to add %v to stack %v\n", string(character), j/step+1)

				push(&stacks[j/step], string(character))

			}
		}
	}

	// Loop through all the stacks :
	for i := 0; i < len(stacks); i++ {

		// Loop through the contents of the stack
		for j := 0; j < len(stacks[i].elements); j++ {
			fmt.Printf("At stack %v there element : %v\n", i, stacks[i].elements[j])
		}
	}

}

func push(stackToUpdate *stack, value string) {
	stackToUpdate.elements = append(stackToUpdate.elements, value)
}

func pop(stackToUpdate *stack) string {

	if len(stackToUpdate.elements) == 0 {
		println("There is nothing to pop")
		return ""
	}

	stackLength := len(stackToUpdate.elements)
	stringToReturn := stackToUpdate.elements[stackLength-1]

	stackToUpdate.elements = stackToUpdate.elements[:stackLength-1]

	return stringToReturn
}
