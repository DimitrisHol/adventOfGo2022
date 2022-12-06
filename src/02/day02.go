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

func part1(data []string) int {

	var totalScore int = 0
	scoringPairs := map[string]int{
		"A X": 4, "A Y": 8, "A Z": 3,
		"B X": 1, "B Y": 5, "B Z": 9,
		"C X": 7, "C Y": 2, "C Z": 6}

	for _, pair := range data {
		totalScore += scoringPairs[pair]
	}
	return totalScore
}

func part2(data []string) int {

	var totalScore int = 0
	scoringPairs := map[string]int{
		"A X": 3, "A Y": 4, "A Z": 8,
		"B X": 1, "B Y": 5, "B Z": 9,
		"C X": 2, "C Y": 6, "C Z": 7}

	for _, pair := range data {
		totalScore += scoringPairs[pair]
	}
	return totalScore
}

func main() {
	data := parseFile("../../input/day02.txt")
	fmt.Printf("Part 1 Score : %v\n", part1(data))
	fmt.Printf("Part 2 Score : %v\n", part2(data))
}
