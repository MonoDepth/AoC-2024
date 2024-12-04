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
		direction := -1
		safe := true
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

			if (diff > 0 && direction < 0) || (diff < 0 && direction > 0) {
				fmt.Println("Unsafe", line)
				safe = false
				break
			}

			if math.Abs(float64(diff)) < 1 || math.Abs(float64(diff)) > 3 {
				fmt.Println("Unsafe", line)
				safe = false
			}
		}

		if safe {
			fmt.Println("Safe", line)
			safeCnt++
		}
	}

	fmt.Println("Safe count", safeCnt)
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
