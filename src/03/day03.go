package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

func parseFile(filename string) []string {

	data, _ := os.ReadFile(filename)
	stringData := string(data)
	return strings.Split(stringData, "\r\n")
}

// [1, 3, 7]
// [6, 7, 10]

// Start with the 1, compare with 7.
// If it is the same, then we found it. Stop everything.

// Otherwise, compare the numbers. If the source number is lesser than
// the destination number, then skip the source number

// If source array is [1,3,7] and destination array is [6,7,10]
// Checking 1 with 6 should let us know that there is no 1 (ascending)
func findDuplicateIntegers(leftSide []int, rightSide []int) (int, bool) {

	var found bool = false

	for i := 0; i < len(leftSide) && !found; i++ {
		for j := 0; j < len(rightSide) && !found; j++ {
			if leftSide[i] == rightSide[j] {
				return leftSide[i], true
			} else if leftSide[i] < rightSide[j] {
				break // Since these are sorted.
			}

		}
	}
	return 0, false
}

// ASCII
// A : 65, Z : 90
// a : 97, z : 122

// Priorities
// A : 27, Z : 52
// a = 1 , z = 26
func convertToPriority(input byte) int {

	if input >= 65 && input <= 90 {
		return int(input) - 38
	} else if input >= 97 && input <= 122 {
		return int(input) - 96
	} else {
		fmt.Printf("Expected A-Z,a-z character, instead got : %v", input)
		panic("Invalid input")
	}
}

func main() {

	data := parseFile("../../input/day03.txt")

	var prioritiesSum int = 0
	for _, ruckSack := range data {

		half := len(ruckSack) / 2
		leftSide := ruckSack[:half]
		rightSide := ruckSack[half:]

		var leftSideInt []int
		var rightSideInt []int

		for i := 0; i < len(leftSide); i++ {
			leftSideInt = append(leftSideInt, convertToPriority(leftSide[i]))
			rightSideInt = append(rightSideInt, convertToPriority(rightSide[i]))
		}

		// Find duplicate numbers between two arrays (sort = fast)
		sort.Ints(leftSideInt)
		sort.Ints(rightSideInt)

		if duplicateValue, ok := findDuplicateIntegers(leftSideInt, rightSideInt); ok {
			prioritiesSum += duplicateValue
		}
	}
	fmt.Printf("Part 1 : PrioritiesSum: %v\n", prioritiesSum)

	var groupPrioritiesSum int = 0
	// Main loop all over again
	for i := 0; i < len(data); i += 3 {

		// This is absolutely horrible I know
		integerCountSack1 := make(map[int]int)
		integerCountSack2 := make(map[int]int)
		integerCountSack3 := make(map[int]int)

		for j := 0; j < len(data[i]); j++ {

			priorityValueForCharacter := convertToPriority(data[i][j])
			integerCountSack1[priorityValueForCharacter] += 1
		}

		for j := 0; j < len(data[i+1]); j++ {

			priorityValueForCharacter := convertToPriority(data[i+1][j])
			integerCountSack2[priorityValueForCharacter] += 1
		}

		for j := 0; j < len(data[i+2]); j++ {

			priorityValueForCharacter := convertToPriority(data[i+2][j])
			integerCountSack3[priorityValueForCharacter] += 1
		}

		// I regret nothing
		for i := 1; i < 53; i++ {
			_, found1 := integerCountSack1[i]
			_, found2 := integerCountSack2[i]
			_, found3 := integerCountSack3[i]

			if found1 && found2 && found3 {

				groupPrioritiesSum += i
				fmt.Printf("i: %v\n", i)
			}

		}
	}

	fmt.Printf("groupPriorities: %v\n", groupPrioritiesSum)

}
