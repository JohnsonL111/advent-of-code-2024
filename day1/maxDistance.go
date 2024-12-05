package main

import (
    "os"
    "bufio"
    "fmt"
	"strconv"
    "strings"
	"sort"
	"math"
)
// helful: https://www.jeremymorgan.com/tutorials/go/how-to-read-text-file-go/
func main() {
	file, ferr := os.Open("input.txt")

	if ferr != nil {
		panic(ferr)
	}
	scanner := bufio.NewScanner(file)
	// slices are dynamically sized arrays
	var list1 []int
	var list2 []int

	// read line by line
	for scanner.Scan() {
		line := scanner.Text()
		items := strings.Fields(line) // account for multiple spaces between lists
		fmt.Printf("list1: %s list2: %s\n", items[0],items[1])

        // Convert strings to integers
        val1, err1 := strconv.Atoi(items[0])
        if err1 != nil {
            fmt.Printf("Error converting %s to int: %v\n", items[0], err1)
            continue
        }

        val2, err2 := strconv.Atoi(items[1])
        if err2 != nil {
            fmt.Printf("Error converting %s to int: %v\n", items[1], err2)
            continue
        }

        // Append to slices
        list1 = append(list1, val1)
        list2 = append(list2, val2)
	}

	fmt.Println(list1)
	fmt.Println(list2)
	// sort
	sort.Ints(list1)
	sort.Ints(list2)
	var total int = 0

	// calculate list
	for i := 0; i < len(list1); i++ {
		diff := math.Abs(float64(list1[i] - list2[i])) // expects float
		total += int(diff)
	}

	fmt.Println("total is", total)
	// return total
}