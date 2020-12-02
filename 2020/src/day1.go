package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func day1() {
	file, err := os.OpenFile("./2020/input/day1.txt", os.O_RDONLY, os.ModePerm)

	if err != nil {
		log.Fatalf("Open file error: %v", err)
		return
	}

	defer file.Close()

	reader := bufio.NewReader(file)
	entries := []int{}

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
		number, convError := strconv.Atoi(trimmed)

		if convError != nil {
			log.Fatalf("Conversion error: %v", convError)
			continue
		}

		entries = append(entries, number)
	}

	log.Printf("Input: %v \n", entries)

	find(entries)
}

func find(entries []int) {
	i := 0
	length := len(entries)

	for i < length {
		entryI := entries[i]
		log.Printf("I %v", entryI)

		j := i + 1

		for j < length {
			entryJ := entries[j]
			log.Printf("J %v", entryJ)

			if (entryJ + entryI >= 2020) {
				j += 1
				continue
			}

			k := j + 1

			for k < length {
				entryK := entries[k]
				log.Printf("K %v", entryK)
				
				if (entryI + entryJ + entryK == 2020) {
					log.Printf("Result %v", entryI * entryJ * entryK)
					return
				}

				k += 1
			}

			j += 1
		}

		i += 1
	}
}
