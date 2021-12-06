package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func Day6() {
	file, err := os.OpenFile("../input/day6.txt", os.O_RDONLY, os.ModePerm)

	if err != nil {
		log.Fatalf("Open file error: %v", err)
		return
	}

	defer file.Close()

	reader := bufio.NewReader(file)
	entries := []string{}

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

		if trimmed != "" {
			entries = append(entries, strings.Split(trimmed, ",")...)
		}
	}

	initial := createLaternFishes(entries)
	final := getFishesAfterDays(initial, 256)
	count := 0

	for _, value := range final {
		count += value
	}

	log.Printf("Total Count %d", count)
}

func createLaternFishes(entries []string) map[int]int {
	laternFishes := make(map[int]int)
	i := 0

	for i < len(entries) {
		key, err := strconv.Atoi(entries[i])

		if err != nil {
			log.Fatal("Cannot be converted to string")
		}

		value, exists := laternFishes[key]

		if exists {
			laternFishes[key] = value + 1
		} else {
			laternFishes[key] = 1
		}

		i += 1
	}

	return laternFishes
}

func copyMap(dict map[int]int) map[int]int {
	copied := make(map[int]int)

	for k, v := range dict {
		copied[k] = v
	}

	return copied
}

func getFishesAfterDays(fishes map[int]int, days int) map[int]int {
	i := 0

	for i < days {
		newMap := make(map[int]int)
		newMap[6] = 0

		for key, value := range fishes {
			key -= 1

			if key == -1 {
				newMap[8] = value
				existing := newMap[6]
				newMap[6] = existing + value
			} else {
				if key == 6 {
					existing := newMap[6]
					value = existing + value
				}

				newMap[key] = value
			}
		}

		fishes = copyMap(newMap)
		i += 1
	}

	return fishes
}
