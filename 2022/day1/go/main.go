package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
)

func panicIf(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	if len(os.Args) < 2 {
		panic("Input file argument required")
	}
	filename := os.Args[1]
	fh, err := os.Open(filename)
	panicIf(err)
	defer fh.Close()

	fmt.Printf("Part 1: %d\n", largestSum(fh, 1))

	_, err = fh.Seek(0, io.SeekStart)
	panicIf(err)

	fmt.Printf("Part 2: %d\n", largestSum(fh, 3))
}

func largestSum(fh io.Reader, nLargest int) int {
	scanner := bufio.NewScanner(fh)
	largestSums := make([]int, 0)
	currentSum := 0
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			largestSums = append(largestSums, currentSum)
			currentSum = 0
		} else {
			currentVal, err := strconv.Atoi(line)
			panicIf(err)
			currentSum += currentVal
		}
	}

	largestSums = append(largestSums, currentSum)

	sort.Ints(largestSums)

	lastNSums := largestSums[len(largestSums)-nLargest:]

	totalSum := 0
	for _, val := range lastNSums {
		totalSum += val
	}

	return totalSum
}
