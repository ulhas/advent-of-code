package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"strings"
)

func day3() {
	a := treeCount(1, 1)
	b := treeCount(3, 1)
	c := treeCount(5, 1)
	d := treeCount(7, 1)
	e := treeCount(1, 2)

	log.Printf("Count %v", a*b*c*d*e)
}

func isTree(inputString string, position int) bool {
	chars := strings.Split(inputString, "")
	return chars[position] == "#"
}

func treeCount(right int, down int) int {
	file, err := os.OpenFile("./2020/input/day3.txt", os.O_RDONLY, os.ModePerm)

	if err != nil {
		log.Fatalf("Open file error: %v", err)
		return -1
	}

	defer file.Close()

	reader := bufio.NewReader(file)

	count := 0
	index := 0
	currentColumn := 0

	for {
		line, err := reader.ReadString('\n')

		if err != nil {
			if err == io.EOF {
				break
			}

			log.Fatalf("Read file line error: %v", err)
			return -1
		}

		trimmed := strings.TrimSpace(line)
		input := trimmed

		for i := 0; i < right*16; i++ {
			trimmed = trimmed + input
		}

		if index == 0 {
			index++
			currentColumn += right
			continue
		}

		if down != 1 && index%down != 0 {
			index++
			continue
		}

		if isTree(trimmed, currentColumn) {
			log.Printf("Tree found at (%v, %v)", currentColumn, index)
			count++
		} else {
			log.Printf("No Tree found at (%v, %v)", currentColumn, index)
		}

		index++
		currentColumn += right
	}

	log.Printf("Count: %v \n", count)
	return count
}
