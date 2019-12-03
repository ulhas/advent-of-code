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

type Coordinate struct {
	x, y int
}

type CoordinateSet struct {
	coordinates []Coordinate
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
	traversed := []Coordinate{}
	minManhattanDistance := 35000

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
		current := Coordinate{0, 0}

		for _, step := range path {
			direction := getDirection(step)
			distance := getDistance(step)

			for j := 1; j <= distance; j++ {
				next := getNextCoordinate(current, direction)

				if index == 0 {
					traversed = append(traversed, next)
				} else {
					if contains(traversed, next) {
						manhattanDistance := getManhattanDistance(next)
						minManhattanDistance = min(manhattanDistance, minManhattanDistance)
					}
				}

				current = next
			}
		}

		index++
	}

	fmt.Printf("Min Manhattan Distance :%v", minManhattanDistance)
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

func contains(coordinates []Coordinate, coordinate Coordinate) bool {
	for _, c := range coordinates {
		if c.x == coordinate.x && c.y == coordinate.y {
			return true
		}
	}

	return false
}

func getNextCoordinate(current Coordinate, direction string) Coordinate {
	next := Coordinate{current.x, current.y}

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
