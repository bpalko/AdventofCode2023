package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var englishToNumber = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	calibrationList := []string{}
	for scanner.Scan() {
		curLine := scanner.Text()
		// regex scan the string for only number, preserving order
		re := regexp.MustCompile(`(?:[1-9]|one|two|three|four|five|six|seven|eight|nine)`) //(`\d`)
		fmt.Println(re.FindAllString(curLine, -1))
		nums := (re.FindAllString(curLine, -1))
		fmt.Printf("nums: %v\n", nums)
		// take the result and grab index[0] and index[-1]
		first := nums[0]
		last := nums[len(nums)-1]
		fmt.Printf("first and last : %v + %v\n", first, last)
		first = processString(first)
		last = processString(last)
		// concat values
		cat := first + last
		// append to master list of calibrations
		calibrationList = append(calibrationList, cat)
	}
	fmt.Printf("calibrationList: %v\n", calibrationList)
	sums := 0
	// sum values within calibrations for final result
	for _, val := range calibrationList {
		valInt, _ := strconv.Atoi(val)
		sums = valInt + sums
	}
	fmt.Printf("sums: %v\n", sums)
}

func processString(originalString string) string {
	substring := strings.ToLower(originalString) // Convert to lowercase for case-insensitive matching

	// Check if the substring is an English representation in the map
	if numberValue, exists := englishToNumber[substring]; exists {
		fmt.Printf("numberValue: %v\n", numberValue)
		return strconv.Itoa(numberValue)
	}
	return originalString
}
