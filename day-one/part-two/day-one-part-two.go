/*
* this task requires us to find the first and last "digit" on each line
* which could be either an actual digit or the word for a digit
* We need to ensure that the code correctly identifies these digits even
* when they are embedded within other letter or at the beginning or end of a line.
*
* approach:
* 1. create a map of number words to their corresponding digits
* 2. define a function to find he first and last digit in each line
* 3. use this function to process each line of the input file, extract the first and last digit, and add them to a sum
* 4. handle any edge cases <non-digit characters>
 */

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// map of spelled out numbers to their digit form
var numberMap = map[string]int{
	"one": 1, "two": 2, "three": 3, "four": 4,
	"five": 5, "six": 6, "seven": 7, "eight": 8, "nine": 9,
}

// check if a character is a digit
func isDigit(ch byte) bool {
	return ch >= '0' && ch <= '9'
}

// update the first and last digit values
func updateDigits(firstDigit *int, lastDigit *int, value int) {
	if *firstDigit == -1 {
		*firstDigit = value
	}
	*lastDigit = value
}

// calculate the sum using the sliding window technique
func calcSum(lines []string) int {
	maxWindowLength := 5 // max length of a spelled-out number
	sum := 0

	for _, line := range lines {
		firstDigit, lastDigit := -1, -1

		for start, end := 0, 0; start < len(line); start++ {
			end = start

			if isDigit(line[start]) {
				digit, _ := strconv.Atoi(string(line[start]))
				updateDigits(&firstDigit, &lastDigit, digit)
				continue
			}

			for end < len(line) && end-start < maxWindowLength {
				if num, ok := numberMap[line[start:end+1]]; ok {
					updateDigits(&firstDigit, &lastDigit, num)
					break
				}
				end++
			}
		}

		if firstDigit != -1 && lastDigit != -1 {
			sum += firstDigit*10 + lastDigit
		}
	}

	return sum

}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	sum := calcSum(lines)
	fmt.Println("Total sum:", sum)

}
