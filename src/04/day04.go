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

func main() {

	data := parseFile("../../input/day04.txt")
	var completeOverlaps int = 0
	var partialOverlaps int = 0

	for _, assignment := range data {

		pair := strings.Split(assignment, ",")

		// ["2-3", "1-4"]
		firstSectorIds := strings.Split(pair[0], "-")
		secondSectorIds := strings.Split(pair[1], "-")

		// ["2", "3"] , ["1" , "4"]
		firstStart, _ := strconv.Atoi(firstSectorIds[0])
		firstEnd, _ := strconv.Atoi(firstSectorIds[1])

		secondStart, _ := strconv.Atoi(secondSectorIds[0])
		secondEnd, _ := strconv.Atoi(secondSectorIds[1])

		// Part 1 :
		if (firstStart >= secondStart && firstEnd <= secondEnd) ||
			(secondStart >= firstStart && secondEnd <= firstEnd) {
			completeOverlaps += 1
		}

		// Part 2 :
		if firstStart <= secondEnd && firstEnd >= secondStart {
			partialOverlaps += 1
		}

	}

	fmt.Printf("Part 1, Complete Overlaps : %v\n", completeOverlaps)
	fmt.Printf("Part 2, Partial Overlaps: %v\n", partialOverlaps)
}
