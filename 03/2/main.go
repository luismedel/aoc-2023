package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, scanner.Err()
}

func toInt(s string) int {
	n, _ := strconv.Atoi(s)
	return n
}

type partNumber struct {
	Row      int
	Position []int
	Number   int
	Valid    bool
}

func getGearRatio(partNumbers []partNumber, pos []int) int {
	row := pos[0]
	col := pos[1]

	count := 0
	total := 1

	for i := range partNumbers {
		pn := &partNumbers[i]
		if pn.Row == row-1 || pn.Row == row+1 {
			if col >= pn.Position[0] && col <= pn.Position[1] {
				total *= pn.Number
				count += 1
			} else if pn.Position[0] == col+1 || pn.Position[1] == col-1 {
				total *= pn.Number
				count += 1
			}
		} else if pn.Row == row {
			if pn.Position[1] == col-1 || pn.Position[0] == col+1 {
				total *= pn.Number
				count += 1
			}
		}
	}

	if count < 2 {
		return 0
	}
	return total
}

func main() {
	lines, _ := readLines("input.txt")

	rpartNumber := regexp.MustCompile(`\d+`)
	rgear := regexp.MustCompile(`\*`)

	sum := 0

	partNumbers := make([]partNumber, 0)
	gearPositions := make([][]int, 0)

	row := 0

	for _, line := range lines {
		numbers := rpartNumber.FindAllStringIndex(line, -1)
		for _, n := range numbers {
			partNumbers = append(partNumbers, partNumber{
				Row:      row,
				Position: []int{n[0], n[1] - 1},
				Number:   toInt(line[n[0]:n[1]]),
				Valid:    false,
			})
		}

		symbols := rgear.FindAllStringIndex(line, -1)
		for _, sym := range symbols {
			gearPositions = append(gearPositions, []int{row, sym[0]})
		}

		row += 1
	}

	for _, symbol := range gearPositions {
		sum += getGearRatio(partNumbers, symbol)
	}

	fmt.Println(sum)
}
