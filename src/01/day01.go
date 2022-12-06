package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func errorHandling(e error) {
	if e != nil {
		panic(e)
	}
}

func parseFile(filename string) []string {

	data, err := os.ReadFile(filename)
	errorHandling(err)
	stringData := string(data)
	return strings.Split(stringData, "\r\n")
}

func main() {

	dataCalories := parseFile("../../input/day01.txt")

	localSums := []int{}
	localSum := 0
	for _, entry := range dataCalories {

		if entry == "" {
			localSums = append(localSums, localSum)
			localSum = 0
		}

		intEntry, _ := strconv.Atoi(entry)
		localSum += intEntry
	}
	localSums = append(localSums, localSum)

	var maxValue int = 0

	for _, calories := range localSums {

		if calories > maxValue {
			maxValue = calories
		}
	}
	// Part 1 : Can you image go doesn't have max() builtin!
	fmt.Printf("The maxValue is: %v\n", maxValue)

	// Part 2 : Dirty... I know
	sort.Ints(localSums)
	result := localSums[len(localSums)-1] + localSums[len(localSums)-2] + localSums[len(localSums)-3]
	fmt.Printf("result: %v\n", result)
}
