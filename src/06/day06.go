package main

import (
	"fmt"
	"os"
	"sort"
)

func parseFile(filename string) string {

	data, _ := os.ReadFile(filename)
	return string(data)
}

func main() {

	data := parseFile("../../input/day06.txt")

	var WINDOW_SIZE_PART_1 int = 4 - 1
	var WINDOW_SIZE_PART_2 int = 14 - 1

	fmt.Printf("Part 1 result : %v\n", calculate(WINDOW_SIZE_PART_1, data))
	fmt.Printf("Part 2 result : %v\n", calculate(WINDOW_SIZE_PART_2, data))

}

func calculate(windowSize int, data string) int {

	for i := windowSize; i < len(data)-4; i++ {

		var windowIsTheOne bool = true

		// 1. Grab the window : 4 or 14
		window := data[i-windowSize : i+1]

		// 2. Sort the window
		windowSorted := []rune(window)
		sort.Slice(windowSorted, func(i int, j int) bool { return windowSorted[i] < windowSorted[j] })

		// 3. Duplications will be adjacent
		for j := 0; j < len(windowSorted)-1; j++ {

			if windowSorted[j] == windowSorted[j+1] {
				windowIsTheOne = false
				break
			}
		}
		if windowIsTheOne {
			return i + 1
		}
	}
	return 0

}
