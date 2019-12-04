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
	file, err := os.OpenFile("./input/2019/day1.txt", os.O_RDONLY, os.ModePerm)

	if err != nil {
		log.Fatalf("Open file error: %v", err)
		return
	}

	defer file.Close()

	reader := bufio.NewReader(file)
	totalFuel := 0

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

		totalFuel += getFuel(number)
	}

	log.Println("********** Day 1 **********")
	log.Printf("Total Fuel equired: %v \n", totalFuel)
	log.Println("********** End Day 1 **********")
}

func getFuel(mass int) int {
	fuel := int(mass/3) - 2

	if fuel <= 0 {
		return 0
	}

	fuel += getFuel(fuel)
	return fuel
}
