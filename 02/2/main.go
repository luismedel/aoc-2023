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

func calcGamePower(games [][]int) int {
	r := 0
	g := 0
	b := 0
	for _, game := range games {
		if game[0] > r {
			r = game[0]
		}
		if game[1] > g {
			g = game[1]
		}
		if game[2] > b {
			b = game[2]
		}
	}
	return max(r, 1) * max(g, 1) * max(b, 1)
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
		_, games := parseLine(line)
		sum += calcGamePower(games)
	}

	fmt.Println(sum)
}
