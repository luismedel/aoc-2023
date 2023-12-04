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

func main() {
	lines, _ := readLines("input.txt")

	r := regexp.MustCompile(`\d`)
	sum := 0

	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		numbers := r.FindAllString(line, -1)
		if len(numbers) == 0 {
			continue
		}
		n := (toInt(numbers[0]) * 10) + toInt(numbers[len(numbers)-1])
		sum += n
	}

	fmt.Println(sum)
}
