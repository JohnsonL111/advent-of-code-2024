package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Helper to check if a report is safe
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

// Function to check if a report can be made safe by removing a single level
// wrapper around pt1 isSafe
func isSafeWithDampener(report []int) bool {
	// Check if the report is already safe
	if isSafe(report) {
		return true
	}

	// Try removing each level and check if the report becomes safe
	for i := 0; i < len(report); i++ {
		// Create a new slice with the i-th level removed
		// Go slice syntax [start, end), simialr to python w/o step
		// ... variadic operator similar to spread in js
		modifiedReport := make([]int, 0, len(report)-1) 
		modifiedReport = append(modifiedReport, report[:i]...) // spread start to i-1
		modifiedReport = append(modifiedReport, report[i+1:]...) // spread i+1 to n

		// Check if the modified report is safe
		if isSafe(modifiedReport) {
			return true
		}
	}

	return false // Not safe even with the dampener
}


func main() {
	file, ferr := os.Open("input.txt")
	if ferr != nil {
		panic(ferr)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	safeCount := 0
	reportIndex := 1

	// Read line by line
	for scanner.Scan() {
		line := scanner.Text()

		// Split the line into individual numbers
		parts := strings.Fields(line)
		// make creates a a slice of ints of a certain length
		report := make([]int, len(parts))
		for i, part := range parts {
			num, err := strconv.Atoi(part)
			if err != nil {
				// %q: quoted, %v: default
				fmt.Printf("Error converting %q to integer: %v\n", part, err)
				return
			}
			report[i] = num
		}

		// Check if the report is safe with the dampener
		if isSafeWithDampener(report) {
			fmt.Printf("Report #%d is SAFE: %v\n", reportIndex, report)
			safeCount++
		} else {
			fmt.Printf("Report #%d is UNSAFE: %v\n", reportIndex, report)
		}

		reportIndex++
	}

	fmt.Printf("Total number of safe reports: %d\n", safeCount)
}
