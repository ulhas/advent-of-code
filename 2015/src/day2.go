package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func Day2() {
	file, err := os.OpenFile("../input/day2.txt", os.O_RDONLY, os.ModePerm)

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

	totalWrappingPaper := getTotalWrappingPaper(entries)
	totalRibbon := getTotalRibbon(entries)

	log.Printf("total wrapping paper %v", totalWrappingPaper)
	log.Printf("total ribbon %v", totalRibbon)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func getTotalWrappingPaper(entries []string) int {
	i := 0
	length := len(entries)
	totalWrappingPaper := 0

	for i < length {
		components := strings.Split(entries[i], "x")
		l, err1 := strconv.Atoi(components[0])
		w, err2 := strconv.Atoi(components[1])
		h, err3 := strconv.Atoi(components[2])

		if err1 != nil || err2 != nil || err3 != nil {
			log.Fatalf("Error in converting")
			return -1
		}

		min := min(min(l*w, w*h), h*l)

		surfaceArea := 2*l*w + 2*w*h + 2*h*l
		wrappingArea := surfaceArea + min

		totalWrappingPaper += wrappingArea

		i += 1
	}

	return totalWrappingPaper
}

func getTotalRibbon(entries []string) int {
	i := 0
	length := len(entries)
	totalRibbon := 0

	for i < length {
		components := strings.Split(entries[i], "x")
		l, err1 := strconv.Atoi(components[0])
		w, err2 := strconv.Atoi(components[1])
		h, err3 := strconv.Atoi(components[2])

		if err1 != nil || err2 != nil || err3 != nil {
			log.Fatalf("Error in converting")
			return -1
		}

		smallestPerimeter := 2 * (l + w + h - (max(max(l, w), h)))
		bow := l * w * h

		totalRibbon += smallestPerimeter + bow

		i += 1
	}

	return totalRibbon
}
