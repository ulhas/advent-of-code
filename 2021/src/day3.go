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

func Day3() {
	file, err := os.OpenFile("../input/day3.txt", os.O_RDONLY, os.ModePerm)

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

	bitCount := getBitCount(entries)
	log.Print(bitCount)
	solveBitCount(len(entries), bitCount)
}

func getBitCount(entries []string) [12]int {
	i := 0
	length := len(entries)
	bitCount := [12]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	for i < length {
		entry := entries[i]
		j := 0
		bits := strings.Split(entry, "")

		for j < 12 {
			char := bits[j]

			if char == "1" {
				number := bitCount[j]
				bitCount[j] = number + 1
			}

			j += 1
		}

		i += 1
	}

	return bitCount
}

func solveBitCount(length int, bitCount [12]int) {
	mostCommonBits := [12]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	leastCommonBits := [12]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	i := 0

	for i < 12 {
		isMost := bitCount[i] > length/2
		if isMost {
			mostCommonBits[i] = 1
		} else {
			leastCommonBits[i] = 1
		}
		i += 1
	}

	log.Print(mostCommonBits)
	log.Print(leastCommonBits)

	gammaRateBit := ""
	epsilonRateBit := ""
	i = 0

	for i < 12 {
		gammaRateBit += fmt.Sprint(mostCommonBits[i])
		epsilonRateBit += fmt.Sprint(leastCommonBits[i])
		i += 1
	}

	gammaRate, err := strconv.ParseInt(gammaRateBit, 2, 16)

	if err != nil {
		log.Fatal("Gamma rate error")
	}

	epsilonRate, err := strconv.ParseInt(epsilonRateBit, 2, 16)

	if err != nil {
		log.Fatal("Epsilon rate error")
	}

	log.Print(gammaRate * epsilonRate)
}
