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

func Day5() {
	file, err := os.OpenFile("../input/day5.txt", os.O_RDONLY, os.ModePerm)

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
			entries = append(entries, trimmed)
		}
	}

	lines := parseLines(entries)
	count := 0

	for _, value := range lines {
		if value > 1 {
			count += 1
		}
	}

	log.Printf("Count %d", count)
}

func parseLines(entries []string) map[string]int {
	lines := make(map[string]int)
	i := 0

	for i < len(entries) {
		entry := entries[i]
		pts := strings.Split(entry, " -> ")
		startPoints := strings.Split(pts[0], ",")
		endPoints := strings.Split(pts[1], ",")

		x1, err1 := strconv.Atoi(startPoints[0])
		y1, err2 := strconv.Atoi(startPoints[1])
		x2, err3 := strconv.Atoi(endPoints[0])
		y2, err4 := strconv.Atoi(endPoints[1])

		if err1 != nil || err2 != nil || err3 != nil || err4 != nil {
			log.Fatal("Cannot convert number")
		}

		points := []string{}

		if x1 == x2 {
			points = getVerticalPoints(y1, y2, x1)
		} else if y1 == y2 {
			points = getHorizontalPoints(x1, x2, y1)
		} else {
			m := getSlope(x1, y1, x2, y2)

			if m == 1 || m == -1 {
				points = getDiagonalPoints(x1, y1, x2, y2, m)
			}
		}

		j := 0

		for j < len(points) {
			point := points[j]
			value, exist := lines[point]

			if exist {
				lines[point] = value + 1
			} else {
				lines[point] = 1
			}

			j += 1
		}

		i += 1
	}

	return lines
}

func getSlope(x1 int, y1 int, x2 int, y2 int) int {
	return (y2 - y1) / (x2 - x1)
}

func min(x int, y int) int {
	if x < y {
		return x
	}

	return y
}

func max(x int, y int) int {
	if x > y {
		return x
	}

	return y
}

func getHorizontalPoints(x1 int, x2 int, y int) []string {
	lines := []string{}
	minX := min(x1, x2)
	maxX := max(x1, x2)

	for minX <= maxX {
		lines = append(lines, fmt.Sprintf("%d,%d", minX, y))
		minX += 1
	}
	return lines
}

func getVerticalPoints(y1 int, y2 int, x int) []string {
	lines := []string{}
	minY := min(y1, y2)
	maxY := max(y1, y2)

	for minY <= maxY {
		lines = append(lines, fmt.Sprintf("%d,%d", x, minY))
		minY += 1
	}
	return lines
}

func getDiagonalPoints(x1 int, y1 int, x2 int, y2 int, m int) []string {
	lines := []string{}
	minX := min(x1, x2)
	maxX := max(x1, x2)
	c := y1 - x1*m

	for minX <= maxX {
		lines = append(lines, fmt.Sprintf("%d,%d", minX, minX*m+c))
		minX += 1
	}

	return lines
}
