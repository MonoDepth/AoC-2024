package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	xmas   = "XMAS"
	xmas_r = "SAMX"
)

func main() {
	lines, err := readFile("input.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}

	total := 0

	for hIdx, line := range lines {
		for vIdx := 0; vIdx < len(line); vIdx++ {

			if len(line) >= vIdx+4 {
				if line[vIdx:vIdx+4] == xmas {
					total++
				}

				if line[vIdx:vIdx+4] == xmas_r {
					total++
				}
			}

			horizontal := extractH(lines, hIdx, vIdx, 4)

			if horizontal == xmas || horizontal == xmas_r {
				total++
			}

			diagonal := extractDiagonal(lines, hIdx, vIdx, 4)
			if diagonal == xmas || diagonal == xmas_r {
				total++
			}

			diagonal_r := extractDiagonal_R(lines, hIdx, vIdx, 4)
			if diagonal_r == xmas || diagonal_r == xmas_r {
				total++
			}
		}
	}

	fmt.Println("Total", total)
}

func extractH(lines []string, hIdx, vIdx, length int) string {
	str := ""
	for i := hIdx; i < len(lines) && len(str) < length; i++ {

		if vIdx >= len(lines[i]) {
			break
		}

		str += string(lines[i][vIdx])
	}
	return str
}

func extractDiagonal(lines []string, hIdx, vIdx, length int) string {
	str := ""
	for i := 0; len(str) < length; i++ {
		if hIdx+i >= len(lines) {
			break
		}

		if vIdx+i >= len(lines[hIdx+i]) {
			break
		}

		str += string(lines[hIdx+i][vIdx+i])
	}
	return str
}

func extractDiagonal_R(lines []string, hIdx, vIdx, length int) string {
	str := ""
	for i := 0; len(str) < length; i++ {
		if hIdx-i < 0 {
			break
		}

		if vIdx+i >= len(lines[hIdx-i]) {
			break
		}

		str += string(lines[hIdx-i][vIdx+i])
	}
	return str
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
