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
	var numberSafe = 0

	err := utils.ReadFileLineByLine(filename, func(line string) {
		isReportSafe, err := isReportSafe(line)
		if err != nil {
			fmt.Println("Error:", err)
		}
		if isReportSafe {
			numberSafe++
		}
	})

	if err != nil {
		fmt.Println("Error:", err)
	}

	fmt.Printf("Number of Safe Reports: %v", numberSafe)
}

func isReportSafe(line string) (bool, error) {
	var increasing bool
	levels := strings.Fields(line)
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
