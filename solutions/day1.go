package solutions

import (
	"fmt"
	"net/http"

	"mikeknowl.es/adventofcode2020go/util"
)

// Day1 responds with the answers for Day 1
func Day1(w http.ResponseWriter, r *http.Request) {
	answer1, err := part1("input/day1.txt")
	answer2, err := part2("input/day1.txt")

	if err != nil {
		w.Write([]byte(fmt.Sprintf("An error occurred: %s\n", err)))
	}

	w.Write([]byte(fmt.Sprintf("Part 1 Answer: %d\nPart 2 Answer: %d", answer1, answer2)))
}

func part1(fileName string) (answer int, err error) {
	ints, err := util.GetIntsFromFile(fileName)
	if err != nil {
		return
	}

	// We want to find the two numbers that add up to 2020
	sum := 2020
	numbers := make(map[int]int)

	// Read through each line of the file
	for _, number := range ints {
		// If the number exists as a key in the map return the answer
		if val, ok := numbers[number]; ok {
			answer = val * number
			return
		}

		// Store the difference>number as key>value
		difference := sum - number
		numbers[difference] = number
	}

	err = fmt.Errorf("%s", "Didn't find two entries that sum to 2020")

	return
}

func part2(fileName string) (answer int, err error) {
	ints, err := util.GetIntsFromFile(fileName)
	if err != nil {
		return
	}

	// We want to find the three numbers that add up to 2020
	sum := 2020
	sums := make(map[int][]int)
	for i := 0; i < len(ints); i++ {
		for j := i + 1; j < len(ints); j++ {
			// Add two numbers and subtract them from 2020 to get the difference
			// Store that in sums as difference>numbers as key>value
			two := []int{ints[i], ints[j]}
			difference := sum - two[0] - two[1]
			sums[difference] = two
		}
	}

	// Check the numbers for the key, return the answer
	for _, number := range ints {
		if val, ok := sums[number]; ok {
			answer = number * val[0] * val[1]
			return
		}
	}

	err = fmt.Errorf("%s", "Didn't find three entries that sum to 2020")
	return
}
