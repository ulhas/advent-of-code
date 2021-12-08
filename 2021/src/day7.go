package main

import (
	"bufio"
	"io"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func Day7() {
	file, err := os.OpenFile("../input/day7.txt", os.O_RDONLY, os.ModePerm)

	if err != nil {
		log.Fatalf("Open file error: %v", err)
		return
	}

	defer file.Close()

	reader := bufio.NewReader(file)
	entries := []int{}

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
		components := strings.Split(trimmed, ",")
		i := 0

		for i < len(components) {
			number, err := strconv.Atoi(components[i])

			if err != nil {
				log.Fatal("Not a number")
			}

			entries = append(entries, number)
			i += 1
		}
	}

	log.Print(getLowestFuel(entries))
}

func abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}

func min7(x int, y int) int {
	if x < y {
		return x
	}

	return y
}

func getLowestFuel(entries []int) int {
	sort.Ints(entries)
	length := len(entries)
	i := 0
	lowestFuel := math.MaxInt

	for i < entries[length-1] {
		j := 0
		fuel := 0

		for j < length {
			distance := abs(entries[j] - i)
			totalDistance := (distance * (distance + 1)) / 2
			fuel += totalDistance
			j += 1
		}

		lowestFuel = min7(lowestFuel, fuel)
		i += 1
	}

	return lowestFuel
}
