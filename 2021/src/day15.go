package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"strings"
)

func Day15() {
	file, err := os.OpenFile("../input/day15.txt", os.O_RDONLY, os.ModePerm)

	if err != nil {
		log.Fatalf("Open file error: %v", err)
		return
	}

	defer file.Close()

	reader := bufio.NewReader(file)
	entries := []string{}
	i := 0

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

		i += 1
	}

}
