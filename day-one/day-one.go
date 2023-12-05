package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	start := time.Now() //start the clock! Santa's got a tight schedule.
	file, err := os.Open("input.txt")

	if err != nil {
		fmt.Fprintf(os.Stderr, "error opening file: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	sum := 0

	for scanner.Scan() {
		line := scanner.Text() //get the current line as a string
		firstDigit, lastDigit := getDigits(line)
		value, err := strconv.Atoi(firstDigit + lastDigit) //concat and convert to int

		if err != nil {
			readError := os.Stderr
			fmt.Println("error reading from input", readError)
			os.Exit(1)
		}
		sum = sum + value //add the current value to the sum

	}
	if scanner.Err() != nil {
		fmt.Fprintf(os.Stderr, "Time to check the list twice, an error occured while reading from input: %v\n", err)
	}
	fmt.Println("Total sum of holiday cheer: ", sum) //print final sum

	elapsed := time.Since(start)
	fmt.Printf("Elves worked for: %s Faster than Rudolph on a snowy night!", elapsed)
}

// helper func
func getDigits(s string) (firstDigit, lastDigit string) {
	for _, char := range s {
		if char >= '0' && char <= '9' {
			firstDigit = string(char)
			break // exit loop after finding first digit
		}
	}

	for i := len(s) - 1; i >= 0; i-- {
		if s[i] >= '0' && s[i] <= '9' {
			lastDigit = string(s[i])
			break //exit loop after finding the last digit
		}
	}
	return
}
