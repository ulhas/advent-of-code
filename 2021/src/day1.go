package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func Day1() {
	file, err := os.OpenFile("../input/day1.txt", os.O_RDONLY, os.ModePerm)

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
		number, convError := strconv.Atoi(trimmed)

		if convError != nil {
			log.Fatalf("Conversion error: %v", convError)
			continue
		}

		entries = append(entries, number)
	}

	solve(entries)
}

func solve(entries []int) {
	i := 0
	length := len(entries)
	windowedEntries := []int{}

	for i < length-2 {
		windowedEntries = append(windowedEntries, entries[i]+entries[i+1]+entries[i+2])
		i++
	}

	length = len(windowedEntries)
	i = 1
	count := 0

	for i < length {
		current := windowedEntries[i]
		previous := windowedEntries[i-1]

		if current > previous {
			count++
		}

		i++
	}

	log.Printf("Count of incremented entries: %d", count)
}
