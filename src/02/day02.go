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

func part1Fast(data []string) int {

	var totalScore int = 0
	scoringPairs := map[string]int{
		"A X": 4, "A Y": 8, "A Z": 3,
		"B X": 1, "B Y": 5, "C Z": 6,
		"C X": 7, "C Y": 2, "B Z": 9}

	for _, pair := range data {
		totalScore += scoringPairs[pair]
	}
	return totalScore
}

func part2(data []string) int {

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

	return totalScorePart

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

func main() {

	data := parseFile("../../input/day02.txt")
	fmt.Printf("Part 1 Score : %v\n", part1Fast(data))
	fmt.Printf("Part 2 Score : %v\n", part2(data))

}
