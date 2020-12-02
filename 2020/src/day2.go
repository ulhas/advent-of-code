package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func day2() {
	file, err := os.OpenFile("./2020/input/day2.txt", os.O_RDONLY, os.ModePerm)

	if err != nil {
		log.Fatalf("Open file error: %v", err)
		return
	}

	defer file.Close()

	reader := bufio.NewReader(file)
	count := 0

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
		valid := isValid(trimmed)

		if valid {
			count++
		}
	}

	log.Printf("Count: %v \n", count)
}

func isValid(input string) bool {
	components := strings.Split(input, " ")

	rule := components[0]
	char := strings.TrimSuffix(components[1], ":")
	password := components[2]

	log.Printf("Rule %v \n", rule)
	log.Printf("Char %v \n", char)
	log.Printf("Password %v \n", password)

	ruleComponent := strings.Split(rule, "-")

	minString := ruleComponent[0]
	maxString := ruleComponent[1]

	min, convError := strconv.Atoi(minString)
	log.Printf("Min %v \n", min)

	if convError != nil {
		log.Fatalf("Conversion error: %v", convError)
		return false
	}

	max, convError2 := strconv.Atoi(maxString)
	log.Printf("Max %v \n", max)

	if convError2 != nil {
		log.Fatalf("Conversion error: %v", convError)
		return false
	}

	chars := strings.Split(password, "")
	return (chars[min-1] == char && chars[max-1] != char || chars[min-1] != char && chars[max-1] == char)
}
