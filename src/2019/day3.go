package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

type Coordinate struct {
	x, y, step int
}

func day3() {
	file, err := os.OpenFile("./input/2019/day3.txt", os.O_RDONLY, os.ModePerm)

	if err != nil {
		log.Fatalf("Open file error: %v", err)
		return
	}

	defer file.Close()

	reader := bufio.NewReader(file)
	index := 0
	traversed := make(map[string]int)
	minManhattanDistance := 35000
	minStepNumber := 60000

	for {
		line, err := reader.ReadString('\n')

		if err != nil {
			if err == io.EOF {
				break
			}

			log.Fatalf("Read file line error: %v", err)
			return
		}

		trimmedLine := strings.TrimSpace(line)
		path := strings.Split(trimmedLine, ",")
		current := Coordinate{0, 0, 0}
		stepNumber := 0

		for _, step := range path {
			direction := getDirection(step)
			distance := getDistance(step)

			for j := 1; j <= distance; j++ {
				stepNumber += 1
				next := getNextCoordinate(current, direction, stepNumber)
				key := next.getKey()
				step, found := traversed[key]

				if index == 0 {
					if !found {
						traversed[key] = next.step
					}
				} else {
					if found {
						manhattanDistance := getManhattanDistance(next)
						minManhattanDistance = min(manhattanDistance, minManhattanDistance)
						numberOfStep := step + next.step
						minStepNumber = min(numberOfStep, minStepNumber)
					}
				}

				current = next
			}
		}

		index++
	}

	log.Printf("\n\nMin Manhattan Distance :%v", minManhattanDistance)
	log.Printf("\n\nMin Step Distance :%v", minStepNumber)
	log.Printf("\n\nEnd")
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

func getNextCoordinate(current Coordinate, direction string, step int) Coordinate {
	next := Coordinate{current.x, current.y, step}

	if direction == "R" {
		next.x += 1
	} else if direction == "L" {
		next.x -= 1
	} else if direction == "U" {
		next.y += 1
	} else {
		next.y -= 1
	}

	return next
}

func getDirection(path string) string {
	return path[:1]
}

func getDistance(path string) int {
	distanceString := path[1:]
	distance, err := strconv.Atoi(distanceString)

	if err != nil {
		log.Fatalf("Error converting distance :%v", err)
		return 0
	}

	return distance
}

func getManhattanDistance(coordinate Coordinate) int {
	return abs(coordinate.x) + abs(coordinate.y)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}

func (coordinate Coordinate) getKey() string {
	x := strconv.Itoa(coordinate.x)
	y := strconv.Itoa(coordinate.y)
	return x + "," + y
}
