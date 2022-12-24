package main

import "fmt"

type stack struct {
	elements []string
	top      int
} // TODO : we might need to have the default value as -1

func main() {

	// This should make 4 stacks
	stacks := make([]stack, 4)

	fmt.Printf("stacks: %v\n", stacks)

	push(&stacks[0], "dimitris")
	push(&stacks[0], "Nick")

	fmt.Printf("stacks[0]: %v\n", stacks[0])

	pop(&stacks[0])

	fmt.Printf("stacks[0]: %v\n", stacks[0])

	fmt.Printf("stacks: %v\n", stacks)
}

func push(stackToUpdate *stack, value string) {
	stackToUpdate.elements = append(stackToUpdate.elements, value)
}

func pop(stackToUpdate *stack) {

	stackLength := len(stackToUpdate.elements)
	stackToUpdate.elements = stackToUpdate.elements[:stackLength-1]
}
