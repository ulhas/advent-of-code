package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func Day9() {
	file, err := os.OpenFile("../input/day9.txt", os.O_RDONLY, os.ModePerm)

	if err != nil {
		log.Fatalf("Open file error: %v", err)
		return
	}

	defer file.Close()

	reader := bufio.NewReader(file)
	entries := [][]int{}

	for {
		line, err := reader.ReadString('\n')

		if err != nil {
			if err == io.EOF {
				break
			}

			log.Fatalf("Read file line error: %v", err)
			return
		}

		trimmed := strings.TrimSpace(line)
		numbers := strings.Split(trimmed, "")
		row := []int{}
		i := 0

		for i < len(numbers) {
			number, err := strconv.Atoi(numbers[i])

			if err != nil {
				log.Fatal()
			}

			row = append(row, number)
			i += 1
		}

		entries = append(entries, row)
	}

	log.Print(calculateRisk(entries))
}

func calculateRisk(entries [][]int) int {
	lowestPoints := []int{}
	rowLength := len(entries)

	for i, row := range entries {
		colLength := len(row)

		for j, cell := range row {
			if j < colLength-1 && cell >= entries[i][j+1] {
				continue
			}

			if j > 0 && cell >= entries[i][j-1] {
				continue
			}

			if i > 0 && cell >= entries[i-1][j] {
				continue
			}

			if i < rowLength-1 && cell >= entries[i+1][j] {
				continue
			}

			lowestPoints = append(lowestPoints, cell)
		}
	}

	i := 0
	risk := 0

	for i < len(lowestPoints) {
		risk = risk + lowestPoints[i] + 1
		i += 1
	}

	return risk
}
