package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	lines, err := readFile("input.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}

	safeCnt := 0
	for _, line := range lines {
		levels := strings.Split(line, " ")

		if isSafe(levels) {
			safeCnt++
		} else {
			for i := 0; i < len(levels); i++ {
				l2 := make([]string, len(levels))
				copy(l2, levels)
				l2 = removeAt(l2, i)
				if isSafe(l2) {
					safeCnt++
					break
				}
			}
		}
	}

	fmt.Println("Safe count", safeCnt)
}

func isSafe(levels []string) bool {
	direction := -1
	for i := 0; i < len(levels)-1; i++ {
		currLevel, err := strconv.Atoi(levels[i])

		if err != nil {
			panic(err)
		}

		nextLevel, err := strconv.Atoi(levels[i+1])
		if err != nil {
			panic(err)
		}

		diff := currLevel - nextLevel

		if i == 0 {
			direction = diff
		}

		if !isValid(direction, currLevel, nextLevel) {
			return false
		}
	}

	return true
}

func isValid(direction, first, second int) bool {
	diff := first - second

	wrongDir := (diff > 0 && direction < 0) || (diff < 0 && direction > 0)
	outOfBounds := math.Abs(float64(diff)) < 1 || math.Abs(float64(diff)) > 3
	return !wrongDir && !outOfBounds
}

func readFile(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}

func removeAt[T any](slice []T, index int) []T {

	return append(slice[:index], slice[index+1:]...)

}
