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
}

func main() {

	data := parseFile("../../input/day05.txt")

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

	lastStackLine := strings.Split(strings.TrimSpace(stackData[len(stackData)-1]), "   ")
	numberOfStacks, _ := strconv.Atoi(lastStackLine[len(lastStackLine)-1])

	// Dynamic size list : slice. Create list of n stacks.
	var stacks = make([]stack, numberOfStacks)

	// To properly parse the input
	var step int = 4

	// Parse the stack input in reverse to fill the stacks
	for i := len(stackData) - 2; i >= 0; i-- {
		for j := 1; j < len(stackData[i]); j += step {
			character := stackData[i][j]

			if character != 32 { // whitespace

				// fmt.Printf("Trying to add %v to stack %v\n", string(character), j/step+1)
				push(&stacks[j/step], string(character))

			}
		}
	}

	// Parsing the move data :
	for i := 0; i < len(moveData); i++ {

		result := strings.Split(moveData[i], " ")

		amount, _ := strconv.Atoi(result[1])
		source, _ := strconv.Atoi(result[3])
		target, _ := strconv.Atoi(result[5])

		source -= 1 // correct index
		target -= 1 // correct index

		// // Part 1 : Individual moving
		// for j := 1; j <= amount; j++ {
		// 	stackPop := pop(&stacks[source])
		// 	push(&stacks[target], stackPop)
		// }

		// Part 2 : Multiple items at a time :

		// 1. Pop the items and add them to a list
		var tempSlice []string
		for j := 1; j <= amount; j++ {
			stackPop := pop(&stacks[source])
			tempSlice = append(tempSlice, stackPop)
		}

		// 2. Reading in reverse, push them to the target stack
		for j := len(tempSlice) - 1; j >= 0; j-- {
			push(&stacks[target], tempSlice[j])
		}

	}

	var resultString string = ""

	for i := 0; i < len(stacks); i++ {
		top := len(stacks[i].elements) - 1
		resultString += stacks[i].elements[top]
	}

	fmt.Printf("resultString: %v\n", resultString)

}

func push(stackToUpdate *stack, value string) {
	stackToUpdate.elements = append(stackToUpdate.elements, value)
}

func pop(stackToUpdate *stack) string {

	if len(stackToUpdate.elements) == 0 {
		panic("Trying to pop element from stack but there is nothing left to pop!")
	}

	stackLength := len(stackToUpdate.elements)
	stringToReturn := stackToUpdate.elements[stackLength-1]

	stackToUpdate.elements = stackToUpdate.elements[:stackLength-1]

	return stringToReturn
}
