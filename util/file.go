package util

import (
	"bufio"
	"os"
	"strconv"
)

// GetIntsFromFile gets each line from a file as an int
func GetIntsFromFile(fileName string) (numbers []int, err error) {
	// Open file to read
	f, err := os.Open("input/day1.txt")
	if err != nil {
		return
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		// Convert the string to a number
		number, convErr := strconv.Atoi(scanner.Text())
		if convErr != nil {
			err = convErr
			return
		}

		numbers = append(numbers, number)
	}

	err = scanner.Err()
	if err != nil {
		return
	}

	return
}
