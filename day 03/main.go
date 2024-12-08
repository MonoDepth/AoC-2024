package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var regex = regexp.MustCompile(`mul\(\d+,\d+\)|do\(\)|don't\(\)`)
var active = true
var total = 0

func main() {
	lines, err := readFile("input.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}

	for _, line := range lines {
		instructions := filterInput(line)
		for _, instruction := range instructions {
			executeInstruction(instruction)
		}
	}

	fmt.Println("Total", total)
}

func filterInput(input string) []string {
	return regex.FindAllString(input, -1)
}

func executeInstruction(instruction string) {
	if instruction == "don't()" {
		active = false
		return
	}

	if instruction == "do()" {
		active = true
		return
	}

	if !active {
		return
	}

	fmt.Println("Got instruction", instruction, "with state", active)
	numbers := strings.Split(strings.TrimSuffix(strings.TrimPrefix(instruction, "mul("), ")"), ",")
	lt, _ := strconv.Atoi(numbers[0])
	rt, _ := strconv.Atoi(numbers[1])
	total += lt * rt
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
