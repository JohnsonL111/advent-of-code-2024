package main
import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// helper to check if report is safe
func isSafe(report []int) bool {
	if len(report) < 2 {
		// A report with fewer than 2 levels is trivially safe
		return true
	}

	// Determine trend (increasing or decreasing)
	var trend int // +1 for increasing, -1 for decreasing, 0 for unset
	for i := 1; i < len(report); i++ {
		diff := report[i] - report[i-1]

		// Check the difference is valid (no equality and within range)
		if diff == 0 || diff < -3 || diff > 3 {
			return false
		}

		// Determine the trend if unset
		if trend == 0 {
			if diff > 0 {
				trend = 1 // Increasing
			} else if diff < 0 {
				trend = -1 // Decreasing
			}
		} else {
			// Check for trend violation
			if (trend == 1 && diff < 0) || (trend == -1 && diff > 0) {
				return false
			}
		}
	}

	return true
}

func main() {
	file, ferr := os.Open("input.txt")
	
	if ferr != nil {
		panic(ferr)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	safeCount := 0
	lineNum := 0

	// read line by line
	for scanner.Scan() {
		lineNum += 1
		line := scanner.Text()

		// Split the line into a string of individual numbers
		parts := strings.Fields(line)
		// create int slice of 0s of same length
		report := make([]int, len(parts))
		for i, part := range parts {
			num, err := strconv.Atoi(part)
			if err != nil {
				// %q: quoted output, %v: default
				fmt.Printf("Error converting %q to integer: %v\n", part, err)
				return
			}
			report[i] = num
		}
		fmt.Println(lineNum, report)

		// Check if the report is safe
		if isSafe(report) {
			fmt.Printf("report #%v is safe", lineNum)
			fmt.Println()
			safeCount++
		}
	}
	fmt.Println("")
	fmt.Println("# of safe reports is ", safeCount)
}