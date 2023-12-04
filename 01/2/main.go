package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
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
	n, err := strconv.Atoi(s)
	if err == nil {
		return n
	}
	switch s {
	case "one":
		return 1
	case "two":
		return 2
	case "three":
		return 3
	case "four":
		return 4
	case "five":
		return 5
	case "six":
		return 6
	case "seven":
		return 7
	case "eight":
		return 8
	case "nine":
		return 9
	}

	return -9999999999999999
}

func expandOverlaps(s string) string {
	replaces := map[string]string{
		"oneight":   "oneeight",
		"twone":     "twoone",
		"threeight": "threeeight",
		"fiveight":  "fiveeight",
		"eighthree": "eightthree",
		"eightwo":   "eighttwo",
		"nineight":  "nineeight",
	}

	for {
		expanded := false
		for k, v := range replaces {
			if strings.Contains(s, k) {
				s = strings.Replace(s, k, v, -1)
				expanded = true
			}
		}
		if !expanded {
			break
		}
	}

	return s
}

func main() {
	lines, _ := readLines("input.txt")

	r := regexp.MustCompile(`\d|one|two|three|four|five|six|seven|eight|nine`)
	sum := 0

	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		numbers := r.FindAllString(expandOverlaps(line), -1)
		if numbers == nil {
			continue
		}
		n := (toInt(numbers[0]) * 10) + toInt(numbers[len(numbers)-1])
		sum += n
	}

	fmt.Println(sum)
}
