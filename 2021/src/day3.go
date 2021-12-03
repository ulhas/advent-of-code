package main

import (
	"bufio"
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

	calculateRating(entries)
}

func calculateRating(entries []string) {
	oxygenGeneratorRatingInBinary := getRating(entries, true, 0)
	co2SrubberRatingInBinary := getRating(entries, false, 0)

	oxygenGeneratorRating, err := strconv.ParseInt(oxygenGeneratorRatingInBinary, 2, 16)

	if err != nil {
		log.Fatal("Oxygen generate rating error")
	}

	co2SrubberRating, err := strconv.ParseInt(co2SrubberRatingInBinary, 2, 16)

	if err != nil {
		log.Fatal("Co2 Scrubber Rating")
	}

	log.Print(oxygenGeneratorRating)
	log.Print(co2SrubberRating)

	log.Print(oxygenGeneratorRating * co2SrubberRating)
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

func getRating(entries []string, shouldFilterByMostCommon bool, position int) string {
	length := len(entries)

	if length == 1 {
		return entries[0]
	}

	bitCount := getBitCount(entries)
	bitCountAtPosition := bitCount[position]

	log.Printf("--- %v", shouldFilterByMostCommon)
	log.Printf("BitCount %v length %v", bitCount, length)

	var filterBitAtPosition string

	if shouldFilterByMostCommon {
		if bitCountAtPosition >= length/2 {
			filterBitAtPosition = "1"
		} else {
			filterBitAtPosition = "0"
		}
	} else {
		if bitCountAtPosition < length/2 {
			filterBitAtPosition = "1"
		} else {
			filterBitAtPosition = "0"
		}
	}

	log.Printf("BitCount at %v %v %v", position, bitCountAtPosition, filterBitAtPosition)

	filteredEntries := []string{}
	i := 0

	for i < len(entries) {
		entry := entries[i]

		if string(entry[position]) == filterBitAtPosition {
			filteredEntries = append(filteredEntries, entry)
		}

		i += 1
	}

	return getRating(filteredEntries, shouldFilterByMostCommon, position+1)
}
