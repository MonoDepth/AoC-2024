package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var regex = regexp.MustCompile(`mul\(\d+,\d+\)`)

func main() {
	lines, err := readFile("input.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}

	total := 0
	for _, line := range lines {
		instructions := filterInput(line)
		for _, instruction := range instructions {
			total += executeInstruction(instruction)
		}
	}

	fmt.Println("Total", total)
}

func filterInput(input string) []string {
	return regex.FindAllString(input, -1)
}

func executeInstruction(instruction string) int {
	fmt.Println("Got instruction", instruction)
	numbers := strings.Split(strings.TrimSuffix(strings.TrimPrefix(instruction, "mul("), ")"), ",")
	lt, _ := strconv.Atoi(numbers[0])
	rt, _ := strconv.Atoi(numbers[1])
	return lt * rt
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
