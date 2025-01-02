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

		var increasing bool
		levels := strings.Fields(line)
		fmt.Println(levels)
		//fmt.Println(len(levels))
		for i := 0; i < len(levels); i++ { // does this need to be len(levels)-2 ????
			fmt.Printf("i: %v\n", i)
			if i == 0 && len(levels) >= 2 {
				if levels[1] > levels[0] {
					increasing = true
				}
				continue
			}

			// not sure this helper method really cleaned anything up here...
			// second iteration is much better... multiple return values is crazy
			lastLevelNum, currentLevelNum, err := convertLevelsToInts(levels[i-1], levels[i])

			fmt.Printf("last: %v, current: %v\n", lastLevelNum, currentLevelNum)

			if err != nil {
				fmt.Println("Error:", err)
				return
			}

			if increasing && currentLevelNum > lastLevelNum && currentLevelNum-lastLevelNum <= 3 && currentLevelNum-lastLevelNum >= 1 {
				if i == len(levels)-1 {
					numberSafe++
				}
				continue //so far so good
			} else if !increasing && currentLevelNum < lastLevelNum && lastLevelNum-currentLevelNum <= 3 && lastLevelNum-currentLevelNum >= 1 { // decreasing
				if i == len(levels)-1 {
					numberSafe++
				}
				continue // so far so good
			} else {
				break // this one fails
			}

		}
	})

	if err != nil {
		fmt.Println("Error:", err)
	}

	fmt.Printf("Number of Safe Reports: %v", numberSafe)
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
