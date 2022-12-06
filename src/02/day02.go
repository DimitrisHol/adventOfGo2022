package main

import (
	"fmt"
	"os"
	"strings"
)

func parseFile(filename string) []string {

	data, _ := os.ReadFile(filename)
	stringData := string(data)
	return strings.Split(stringData, "\r\n")
}

func calculatePickPoints(pick string) int {

	if pick == "X" {
		return 1
	} else if pick == "Y" {
		return 2
	} else if pick == "Z" {
		return 3
	} else {
		panic("Things went wrong")
	}
}

func part2(data []string) {

	var totalScorePart int = 0

	for _, pair := range data {
		var localScore int = 0

		picks := strings.Split(pair, " ")

		opponentPick := picks[0]
		strategy := picks[1]

		// Step 1 : Follow the strategy :
		if strategy == "X" { // LOSS :

			if opponentPick == "A" {
				strategy = "Z"
			} else if opponentPick == "B" {
				strategy = "X"
			} else {
				strategy = "Y"
			}

		} else if strategy == "Y" {

			if opponentPick == "A" {
				strategy = "X"
			} else if opponentPick == "B" {
				strategy = "Y"
			} else {
				strategy = "Z"
			}

		} else if strategy == "Z" {

			if opponentPick == "A" {
				strategy = "Y"
			} else if opponentPick == "B" {
				strategy = "Z"
			} else {
				strategy = "X"
			}

		}

		// Step 2 : Bonus points
		pickPoints := calculatePickPoints(strategy)
		localScore += pickPoints

		// Step 2 : Part 1 Win condition Resolution
		var winCondition bool = strategy == "X" && opponentPick == "C" ||
			strategy == "Y" && opponentPick == "A" ||
			strategy == "Z" && opponentPick == "B"

		var drawCondition bool = strategy == "X" && opponentPick == "A" ||
			strategy == "Y" && opponentPick == "B" ||
			strategy == "Z" && opponentPick == "C"

		var lossCondition bool = strategy == "X" && opponentPick == "B" ||
			strategy == "Y" && opponentPick == "C" ||
			strategy == "Z" && opponentPick == "A"

		// Step 3 : Award points :
		if winCondition {
			localScore += 6
		} else if drawCondition {
			localScore += 3
		} else if lossCondition {
		} else {
			panic("Something went wrong")
		}

		totalScorePart += localScore

	}

	fmt.Printf("Part 2: %v\n", totalScorePart)

}

func main() {

	data := parseFile("../../input/day02.txt")

	part2(data)

	os.Exit(0)

	var totalScorePart int = 0

	for _, pair := range data {
		var localScore int = 0

		picks := strings.Split(pair, " ")

		opponentPick := picks[0]
		myPick := picks[1]

		// Step 1 : Bonus points
		pickPoints := calculatePickPoints(myPick)
		localScore += pickPoints

		// Step 2 : Part 1 Win condition Resolution
		var winCondition bool = myPick == "X" && opponentPick == "C" ||
			myPick == "Y" && opponentPick == "A" ||
			myPick == "Z" && opponentPick == "B"

		var drawCondition bool = myPick == "X" && opponentPick == "A" ||
			myPick == "Y" && opponentPick == "B" ||
			myPick == "Z" && opponentPick == "C"

		var lossCondition bool = myPick == "X" && opponentPick == "B" ||
			myPick == "Y" && opponentPick == "C" ||
			myPick == "Z" && opponentPick == "A"

		// Step 3 : Award points :
		if winCondition {
			localScore += 6
		} else if drawCondition {
			localScore += 3
		} else if lossCondition {
		} else {
			panic("Something went wrong")
		}

		totalScorePart += localScore

	}

	fmt.Printf("Part 1 Total Score: %v\n", totalScorePart)

}
