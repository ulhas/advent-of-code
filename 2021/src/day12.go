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
	caves []*Cave
}

func (cave *Cave) addCave(neighbor *Cave) {
	cave.adjacent = append(cave.adjacent, neighbor)
}

func (cave *Cave) isSmall() bool {
	return strings.ToLower(cave.name) == cave.name
}

func (cave *Cave) isStart() bool {
	return cave.name == "start"
}

func (cave *Cave) isEnd() bool {
	return cave.name == "end"
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
	caveMap.print()
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

func (caveMap *Map) print() {
	visited := make(map[string]int)
	path := []string{}
	count := 0
	caveMap.printFrom(caveMap.start, visited, path, &count)
	log.Print(count)
}

func (caveMap *Map) printFrom(cave *Cave, visited map[string]int, path []string, count *int) {
	if cave.isStart() || cave.isEnd() {
		visited[cave.name] = 1
	} else if cave.isSmall() {
		value, exists := visited[cave.name]

		if exists {
			visited[cave.name] = value + 1
		} else {
			visited[cave.name] = 1
		}
	}

	path = append(path, cave.name)

	if cave.name == "end" {
		*count += 1
	} else {
		i := 0

		for i < len(cave.adjacent) {
			adj := cave.adjacent[i]
			_, exists := visited[adj.name]

			if !exists {
				caveMap.printFrom(adj, visited, path, count)
			} else {
				if adj.isStart() || adj.isEnd() || hasTwoVisit(visited) {
					i += 1
					continue
				}

				caveMap.printFrom(adj, visited, path, count)
			}

			i += 1
		}
	}

	path = path[:1]
	value, _ := visited[cave.name]

	if value <= 1 {
		delete(visited, cave.name)
	} else {
		visited[cave.name] = value - 1
	}
}

func hasTwoVisit(visited map[string]int) bool {
	for key, value := range visited {
		if key == "start" || key == "end" {
			continue
		}

		if value == 2 {
			return true
		}
	}
	return false
}
