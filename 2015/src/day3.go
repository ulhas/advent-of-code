package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

type Location struct {
	x, y int
}

func Day3() {
	file, err := os.OpenFile("../input/day3.txt", os.O_RDONLY, os.ModePerm)

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
		components := strings.Split(trimmed, "")
		entries = append(entries, components...)
	}

	locations := getLocations(entries)
	log.Print(len(locations))
}

func getLocations(entries []string) map[string]int {
	visits := map[string]int{"0,0": 2}
	i := 0
	santa := Location{x: 0, y: 0}
	robot := Location{x: 0, y: 0}

	for i < len(entries) {
		entry := entries[i]

		if i%2 == 0 {
			santa = santa.move(entry)
			location := santa.toString()
			value, exists := visits[location]

			if exists {
				visits[location] = value + 1
			} else {
				visits[location] = 1
			}

		} else {
			robot = robot.move(entry)
			location := robot.toString()
			value, exists := visits[location]

			if exists {
				visits[location] = value + 1
			} else {
				visits[location] = 1
			}
		}

		i += 1
	}

	return visits
}

func (location Location) move(entry string) Location {
	if entry == "^" {
		location.y += 1
	} else if entry == "v" {
		location.y -= 1
	} else if entry == ">" {
		location.x += 1
	} else if entry == "<" {
		location.x -= 1
	}

	return Location{x: location.x, y: location.y}
}

func (location Location) toString() string {
	return fmt.Sprintf("%v,%v", location.x, location.y)
}
