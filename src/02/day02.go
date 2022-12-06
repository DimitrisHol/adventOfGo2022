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

func main() {
	data := parseFile("../../input/day02.txt")

	var part1, part2 int = 0, 0

	part1Scores := map[string]int{
		"A X": 4, "A Y": 8, "A Z": 3,
		"B X": 1, "B Y": 5, "B Z": 9,
		"C X": 7, "C Y": 2, "C Z": 6}

	part2Scores := map[string]int{
		"A X": 3, "A Y": 4, "A Z": 8,
		"B X": 1, "B Y": 5, "B Z": 9,
		"C X": 2, "C Y": 6, "C Z": 7}

	for _, pair := range data {
		part1 += part1Scores[pair]
		part2 += part2Scores[pair]
	}

	fmt.Printf("Part 1 Score : %v\n", part1)
	fmt.Printf("Part 2 Score : %v\n", part2)
}
