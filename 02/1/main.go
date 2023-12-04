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
	n, _ := strconv.Atoi(s)
	return n
}

func isValidGame(game []int) bool {
	return game[0] <= 12 && game[1] <= 13 && game[2] <= 14
}

func parseLine(line string) (int, [][]int) {
	rline := regexp.MustCompile(`^Game (\d+)\:\s*(.+)$`)
	rdice := regexp.MustCompile(`(\d+) (red|green|blue)`)

	matches := rline.FindStringSubmatch(line)
	gameId := toInt(matches[1])
	games := make([][]int, 0)
	for _, parts := range strings.Split(matches[2], ";") {
		game := []int{0, 0, 0}

		for _, values := range rdice.FindAllStringSubmatch(parts, -1) {
			switch values[2] {
			case "red":
				game[0] = toInt(values[1])
			case "green":
				game[1] = toInt(values[1])
			case "blue":
				game[2] = toInt(values[1])
			}
		}

		games = append(games, game)
	}

	return gameId, games
}

func main() {
	lines, _ := readLines("input.txt")

	sum := 0

	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		gameId, games := parseLine(line)
		fmt.Println(gameId, games)

		valid := true
		for _, game := range games {
			if !isValidGame(game) {
				valid = false
				break
			}
			if !valid {
				break
			}
		}

		if valid {
			sum += gameId
		}
	}

	fmt.Println(sum)
}
