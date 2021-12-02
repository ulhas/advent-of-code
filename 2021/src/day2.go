package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

type Pair struct {
	direction string
	units     int
}

func Day2() {
	file, err := os.OpenFile("../input/day2.txt", os.O_RDONLY, os.ModePerm)

	if err != nil {
		log.Fatalf("Open file error: %v", err)
		return
	}

	defer file.Close()

	reader := bufio.NewReader(file)
	entries := []Pair{}

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
		components := strings.Split(trimmed, " ")

		number, convError := strconv.Atoi(components[1])

		if convError != nil {
			log.Fatalf("Conversion error: %v", convError)
			continue
		}

		pair := Pair{direction: components[0], units: number}
		entries = append(entries, pair)
	}

	solveDay2(entries)
}

func solveDay2(entries []Pair) {
	i := 0
	length := len(entries)
	horizontal := 0
	depth := 0
	aim := 0

	for i < length {
		entry := entries[i]

		if entry.direction == "forward" {
			horizontal += entry.units
			depth += aim * entry.units
		} else if entry.direction == "up" {
			aim -= entry.units
		} else if entry.direction == "down" {
			aim += entry.units
		}

		i += 1
	}

	log.Printf("Result: %d", horizontal*depth)
}
