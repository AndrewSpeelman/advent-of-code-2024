package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
)

func panicIfNil(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	reader, err := os.Open("input.txt")
	panicIfNil(err)

	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanWords)

	//seperate out values
	var left []int
	var right []int

	counter := 0
	for scanner.Scan() {
		x, err := strconv.Atoi(scanner.Text())
		panicIfNil(err)

		if counter%2 == 0 {
			left = append(left, x)
		} else {
			right = append(right, x)
		}
		counter++
	}

	slices.Sort(left)
	slices.Sort(right)

	var distanceValue float64

	if len(left) == len(right) {
		for i := range left {
			distanceValue = distanceValue + (math.Abs(float64(left[i]) - float64(right[i])))
		}
	}

	// "total distance value"
	fmt.Println(int(distanceValue))

	// now to find "similarity score"
	// (left value ) * (number of occurances in right list)
	var similarityScore int
	for _, value := range left {
		similarityScore = similarityScore + (value * countOccurrences(value, right))
	}

	// "similarity score"
	fmt.Println(similarityScore)
}

func countOccurrences(val int, slice []int) int {
	count := 0
	for _, v := range slice {
		if v == val {
			count++
		}
	}
	return count
}
