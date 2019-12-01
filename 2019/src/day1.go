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

func main() {
	pwd, _ := os.Getwd()
	file, err := os.OpenFile(pwd+"/2019/input/day1.txt", os.O_RDONLY, os.ModePerm)

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

	fmt.Printf("Total Fuel equired: %v", totalFuel)
}

func getFuel(mass int) int {
	fuel := int(mass/3) - 2

	if fuel <= 0 {
		return 0
	}

	fuel += getFuel(fuel)
	return fuel
}
