package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type LowestPoint struct {
	number int
	row    int
	col    int
}

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

	lowestPoints := getLowestPoints(entries)
	log.Printf("Risk %d", calculateRisk(lowestPoints))
	log.Printf("Basin %d", calculateBasin(entries, lowestPoints))
}

func getLowestPoints(entries [][]int) []LowestPoint {
	lowestPoints := []LowestPoint{}
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

			lowestPoints = append(lowestPoints, LowestPoint{number: cell, row: i, col: j})
		}
	}

	return lowestPoints
}

func calculateRisk(lowestPoints []LowestPoint) int {
	i := 0
	risk := 0

	for i < len(lowestPoints) {
		risk = risk + lowestPoints[i].number + 1
		i += 1
	}

	return risk
}

func calculateBasin(entries [][]int, lowestPoints []LowestPoint) int {
	i := 0
	sizes := []int{}

	for i < len(lowestPoints) {
		lowestPoint := lowestPoints[i]
		basin := getBasin(entries, lowestPoint.row, lowestPoint.col, make(map[string]int))
		sizes = append(sizes, len(basin))
		i += 1
	}

	sort.Ints(sizes)
	length := len(sizes)

	return sizes[length-1] * sizes[length-2] * sizes[length-3]
}

func getBasin(entries [][]int, currRow int, currCol int, basin map[string]int) map[string]int {
	if entries[currRow][currCol] == 9 {
		return basin
	}

	key := fmt.Sprintf("%d,%d", currRow, currCol)
	_, exists := basin[key]

	if exists {
		return basin
	}

	basin[key] = entries[currRow][currCol]

	if currRow < len(entries)-1 {
		basin = mergeMaps(basin, getBasin(entries, currRow+1, currCol, basin))
	}

	if currRow > 0 {
		basin = mergeMaps(basin, getBasin(entries, currRow-1, currCol, basin))
	}

	if currCol < len(entries[0])-1 {
		basin = mergeMaps(basin, getBasin(entries, currRow, currCol+1, basin))
	}

	if currCol > 0 {
		basin = mergeMaps(basin, getBasin(entries, currRow, currCol-1, basin))
	}

	return basin
}

func mergeMaps(a map[string]int, b map[string]int) map[string]int {
	for k, v := range b {
		a[k] = v
	}

	return a
}
