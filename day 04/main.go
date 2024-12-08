package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	xmas   = "MAS"
	xmas_r = "SAM"
)

func main() {
	lines, err := readFile("input.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}

	total := 0

	for hIdx, line := range lines {

		if hIdx == 0 || hIdx == len(lines)-1 {
			continue
		}

		for vIdx := 1; vIdx < len(line)-1; vIdx++ {
			t := 0
			if line[vIdx] != 'A' {
				continue
			}

			diag := string(lines[hIdx-1][vIdx-1]) + string(lines[hIdx][vIdx]) + string(lines[hIdx+1][vIdx+1])
			diag_r := string(lines[hIdx+1][vIdx-1]) + string(lines[hIdx][vIdx]) + string(lines[hIdx-1][vIdx+1])

			if diag == xmas || diag == xmas_r {
				t++
			}

			if diag_r == xmas || diag_r == xmas_r {
				t++
			}

			if t == 2 {
				total++
			}
		}
	}
	fmt.Println("Total", total)
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
