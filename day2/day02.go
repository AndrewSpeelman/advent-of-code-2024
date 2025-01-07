package main

import (
	"fmt"
	"main/utils"
	"strconv"
	"strings"
)

func main() {
	filename := "input.txt"

	//So, a report only counts as safe if both of the following are true:
	//The levels are either all increasing or all decreasing.
	//Any two adjacent levels differ by at least one and at most three.

	numberSafe := getNumberSafe(filename)

	fmt.Printf("Number of Safe Reports: %v", numberSafe)
}

func getNumberSafe(filename string) int {
	var numberSafe = 0
	err := utils.ReadFileLineByLine(filename, func(line string) {
		fmt.Println("---------------------------------")
		isReportSafe, err := isReportSafe(line)
		fmt.Println("---------------------------------")

		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		if isReportSafe {
			numberSafe++
		}
	})

	if err != nil {
		fmt.Println("Error:", err)
	}
	return numberSafe
}

func isReportSafe(line string) (bool, error) {
	levels := strings.Fields(line)
	isSafe, err := isLevelsSafe(levels)
	if err != nil {
		return false, nil
	}
	if isSafe {
		return true, nil
	} else {
		// run dampener brute force baby
		for i := 0; i <= len(levels)-1; i++ {
			// remove each level and re-try
			fmt.Printf("i: %v\n", i)
			//copyLevels := append([]int{}, levels[]...)
			//testLevels := append(levels[:i], copyLevels[i+1:]...)
			testLevels := deleteNth(levels, i)
			fmt.Printf("Testing level: %v\n", testLevels)
			isSafe, err := isLevelsSafe(testLevels)
			if err != nil {
				return false, err
			} else if isSafe {
				return true, nil
			}

		}
		return false, nil
	}
}

func isLevelsSafe(levels []string) (bool, error) {
	fmt.Printf("isLevelsSafe called with input: %v\n", levels)
	var increasing bool
	for i := 1; i < len(levels); i++ {
		// not sure this helper method really cleaned anything up here...
		// second iteration is much better... multiple return values is crazy
		lastLevelNum, currentLevelNum, err := convertLevelsToInts(levels[i-1], levels[i])

		if i == 1 && currentLevelNum > lastLevelNum {
			increasing = true
			//fmt.Printf("Level is increasing\n")
		}

		if err != nil {
			fmt.Println("Error:", err)
			return false, err
		}

		if increasing && currentLevelNum > lastLevelNum && currentLevelNum-lastLevelNum <= 3 && currentLevelNum-lastLevelNum >= 1 {
			if i == len(levels)-1 {
				fmt.Printf("Level is correct: %v\n", levels)
				return true, nil
			}
			continue //so far so good
		} else if !increasing && currentLevelNum < lastLevelNum && lastLevelNum-currentLevelNum <= 3 && lastLevelNum-currentLevelNum >= 1 { // decreasing
			if i == len(levels)-1 {
				fmt.Printf("Level is correct: %v\n", levels)
				return true, nil
			}
			continue // so far so good
		} else {
			//fmt.Printf("Level is wrong: %v\n", levels)
			return false, nil
		}
	}
	//fmt.Println("---------------------------------")
	return false, nil
}

//340 too high
//380 wrong

// 249 not right... must be getting close to this value (didn't say too high or low, 2x in a row)
// 483 not right
// 261 not right ... AGHHHHHH

func convertLevelsToInts(lastLevel string, currentLevel string) (int, int, error) {
	lastLevelNum, err := strconv.Atoi(lastLevel)

	// I wonder if this error handling can be condensed into a single block?
	if err != nil {
		fmt.Println("Error failed to parse level into int: "+lastLevel, err)
		return 0, 0, err
	}

	currentLevelNum, err := strconv.Atoi(currentLevel)
	if err != nil {
		fmt.Println("Error failed to parse level into int: "+currentLevel, err)
		return 0, 0, err
	}
	return lastLevelNum, currentLevelNum, nil
}

func deleteNth(slice []string, n int) []string {
	if n < 0 || n >= len(slice) {
		return slice // Return the original slice if the index is out of bounds
	}

	// Create a copy of the slice to ensure original slice is not affected
	newSlice := append([]string(nil), slice...) // Create a copy of the slice

	// Concatenate the elements before and after the nth element
	return append(newSlice[:n], newSlice[n+1:]...)
}
