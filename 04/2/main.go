package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"slices"
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
	n, err := strconv.Atoi(strings.Trim(s, " "))
	if err != nil {
		fmt.Println(s)
		return -999999
	}
	return n
}

func toIntArray(s []string) []int {
	numbers := make([]int, 0)
	for _, n := range s {
		if n == "" {
			continue
		}
		numbers = append(numbers, toInt(n))
	}
	return numbers
}

type cardInfo struct {
	CardId         int
	WinningNumbers []int
	Numbers        []int
}

func (c *cardInfo) calculateCardScore() int {
	var wins int = 0
	for _, n := range c.Numbers {
		if slices.Contains(c.WinningNumbers, n) {
			wins += 1
		}
	}
	return wins
}

func processCard(index int, cards []cardInfo, cached map[int]int) int {
	_, exists := cached[index]
	if exists {
		fmt.Println("Cache hit", index)
		return cached[index]
	}

	var processed int = 1

	score := cards[index].calculateCardScore()
	var copyIndex int = index + 1
	for i := 0; i < score; i++ {
		fmt.Println(index+1, "spawns", copyIndex+1)
		processed += processCard(copyIndex, cards, cached)
		copyIndex = copyIndex + 1
	}

	cached[index] = processed
	return processed
}

func parseLine(line string) cardInfo {
	r := regexp.MustCompile(`^Card\s+(\d+)\: (.*?)\s+\|\s+(.*?)$`)
	matches := r.FindStringSubmatch(line)
	return cardInfo{
		CardId:         toInt(matches[1]),
		WinningNumbers: toIntArray(strings.Split(matches[2], " ")),
		Numbers:        toIntArray(strings.Split(matches[3], " ")),
	}
}

func main() {
	lines, _ := readLines("input.txt")

	sum := 0

	cards := make([]cardInfo, 0)
	for _, line := range lines {
		card := parseLine(line)
		cards = append(cards, card)
	}

	cache := make(map[int]int)
	for i := range cards {
		cardScore := processCard(i, cards, cache)
		sum += cardScore
		fmt.Println(cards[i], cardScore)
	}

	fmt.Println(sum)
}
