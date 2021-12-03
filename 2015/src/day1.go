package main

import (
	"bufio"
	"io"
	"log"
	"os"
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
		entries = append(entries, strings.Split(trimmed, "")...)
	}

	day1(entries)
}

func day1(entries []string) {
	i := 0
	length := len(entries)
	floor := 0

	for i < length {
		if entries[i] == "(" {
			floor += 1
		} else {
			floor -= 1
		}

		if floor == -1 {
			break
		}
		i += 1
	}

	log.Print(i)
}
