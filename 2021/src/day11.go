package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func Day11() {
	file, err := os.OpenFile("../input/day11.txt", os.O_RDONLY, os.ModePerm)

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

	log.Print(getFlashCount(entries))
}

func getFlashCount(entries [][]int) int {
	flash := 0
	step := 0

	for true {
		record := make(map[string]bool)

		for i, row := range entries {
			for j, _ := range row {
				flash += getFlashesForStep(entries, i, j, record)
			}
		}

		if len(record) == 100 {
			break
		}

		step += 1
	}

	log.Print(step + 1)

	return flash
}

func getFlashesForStep(entries [][]int, currRow int, currCol int, record map[string]bool) int {
	flash := 0
	key := fmt.Sprintf("%d,%d", currRow, currCol)
	_, exists := record[key]

	if exists {
		return flash
	}

	value := entries[currRow][currCol]

	if value < 9 {
		entries[currRow][currCol] = value + 1
		return flash
	}

	entries[currRow][currCol] = 0
	record[key] = true
	flash = 1

	if currRow < len(entries)-1 {
		flash += getFlashesForStep(entries, currRow+1, currCol, record)

		if currCol < len(entries[0])-1 {
			flash += getFlashesForStep(entries, currRow+1, currCol+1, record)
		}

		if currCol > 0 {
			flash += getFlashesForStep(entries, currRow+1, currCol-1, record)
		}
	}

	if currRow > 0 {
		flash += getFlashesForStep(entries, currRow-1, currCol, record)

		if currCol < len(entries[0])-1 {
			flash += getFlashesForStep(entries, currRow-1, currCol+1, record)
		}

		if currCol > 0 {
			flash += getFlashesForStep(entries, currRow-1, currCol-1, record)
		}
	}

	if currCol < len(entries[0])-1 {
		flash += getFlashesForStep(entries, currRow, currCol+1, record)
	}

	if currCol > 0 {
		flash += getFlashesForStep(entries, currRow, currCol-1, record)
	}

	return flash
}
