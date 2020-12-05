package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"strings"
)

func day5() {
	file, err := os.OpenFile("./2020/input/day5.txt", os.O_RDONLY, os.ModePerm)

	if err != nil {
		log.Fatalf("Open file error: %v", err)
		return
	}

	defer file.Close()

	reader := bufio.NewReader(file)
	seatIDs := []int{}
	max := 0

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
		seatID := getSeatID(trimmed)

		if seatID > max {
			max = seatID
		}

		seatIDs = append(seatIDs, seatID)
	}

	log.Printf("Max Number %v", max)

	mySeatID := getMySeatID(seatIDs)
	log.Printf("Seat ID %v", mySeatID)
}

func getSeatID(seat string) int {
	components := strings.Split(seat, "")

	minRow := 0
	maxRow := 127
	row := 0

	minCol := 0
	maxCol := 7
	col := 0

	for i := 0; i < len(components); i++ {
		element := components[i]

		if i < 7 {
			delta := (maxRow + 1 - minRow) / 2

			if element == "F" {
				maxRow = maxRow - delta
			} else {
				minRow = minRow + delta
			}

			if i == 6 {
				if components[i] == "B" {
					row = maxRow
				} else {
					row = minRow
				}
			}
		} else {
			delta := (maxCol + 1 - minCol) / 2

			if element == "L" {
				maxCol = maxCol - delta
			} else {
				minCol = minCol + delta
			}

			if i == 9 {
				if components[i] == "L" {
					col = minCol
				} else {
					col = maxCol
				}
			}
		}
	}

	seatID := row*8 + col

	return seatID
}

func getMySeatID(seatIDs []int) int {
	for i := 0; i < len(seatIDs); i++ {
		current := seatIDs[i]

		result1 := findInt(seatIDs, current+1)
		result2 := findInt(seatIDs, current+2)

		if result1 == -1 && result2 != -1 {
			return current + 1
		}
	}

	return -1
}

func findInt(slice []int, number int) int {
	for i := 0; i < len(slice); i++ {
		if number == slice[i] {
			return i
		}
	}

	return -1
}
