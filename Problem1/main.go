package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

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
		re := regexp.MustCompile(`\d`)
		fmt.Println(re.FindAllString(curLine, -1))
		nums := (re.FindAllString(curLine, -1))
		fmt.Printf("nums: %v\n", nums)
		// take the result and grab index[0] and index[-1]
		first := nums[0]
		last := nums[len(nums)-1]
		fmt.Printf("first and last : %v + %v\n", first, last)
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
