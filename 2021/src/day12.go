package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"strings"
)

type Cave struct {
	name     string
	adjacent []*Cave
}

type Map struct {
	start *Cave
}

func (cave *Cave) addCave(neighbor *Cave) {
	cave.adjacent = append(cave.adjacent, neighbor)
}

func (cave *Cave) isSmall() bool {
	return strings.ToLower(cave.name) == cave.name
}

func Day12() {
	file, err := os.OpenFile("../input/day12.txt", os.O_RDONLY, os.ModePerm)

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
		entries = append(entries, trimmed)
	}

	caveMap := createMap(entries)
	log.Print(caveMap)
}

func createMap(entries []string) Map {
	i := 0
	created := make(map[string]*Cave)
	var start *Cave

	for i < len(entries) {
		entry := entries[i]
		caves := strings.Split(entry, "-")
		firstName := caves[0]
		secondName := caves[1]

		firstCave, firstExists := created[firstName]

		if !firstExists {
			firstCave = &Cave{name: firstName, adjacent: []*Cave{}}
			created[firstName] = firstCave

			if firstName == "start" {
				start = firstCave
			}
		}

		secondCave, secondExists := created[secondName]

		if !secondExists {
			secondCave = &Cave{name: secondName, adjacent: []*Cave{}}
			created[secondName] = secondCave

			if secondName == "start" {
				start = secondCave
			}
		}

		firstCave.addCave(secondCave)
		secondCave.addCave(firstCave)

		i += 1
	}

	return Map{start: start}
}
