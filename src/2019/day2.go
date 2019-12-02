package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func day2() {
	file, err := os.OpenFile("./input/2019/day2.txt", os.O_RDONLY, os.ModePerm)

	if err != nil {
		log.Fatalf("Open file error: %v", err)
		return
	}

	defer file.Close()

	reader := bufio.NewReader(file)
	line, err := reader.ReadString('\n')

	if err != nil {
		log.Fatalf("Read file line error: %v", err)
		return
	}

	input := getIntCode(line)
	fmt.Printf("Input IntCode :%v", input)

	output := getCalculatedIntCode(input)
	fmt.Printf("\n\nOutput IntCode :%v", output)
}

func getIntCode(inputLine string) []int {
	trimmedLine := strings.TrimSpace(inputLine)
	inputs := strings.Split(trimmedLine, ",")
	inputNumber := []int{}

	for _, input := range inputs {
		converted, err := strconv.Atoi(input)

		if err != nil {
			log.Fatalf("\n\nError converting input : %v", err)
			continue
		}

		inputNumber = append(inputNumber, converted)
	}

	return inputNumber
}

func getCalculatedIntCode(intCode []int) []int {
	position := 0
	length := len(intCode)

	for position <= length {
		opCode := intCode[position]

		if opCode == 99 {
			fmt.Print("\n\nOpcode 99 encountered. Aborting")
			return intCode
		}

		if opCode == 1 || opCode == 2 {
			firstPosition := intCode[position+1]
			secondPosition := intCode[position+2]
			resultPosition := intCode[position+3]

			if opCode == 1 {
				intCode[resultPosition] = intCode[firstPosition] + intCode[secondPosition]
			} else {
				intCode[resultPosition] = intCode[firstPosition] * intCode[secondPosition]
			}
		} else {
			log.Fatalf("\n\nInvalid OpCode :%v", opCode)
			break
		}

		position += 4
	}

	return intCode
}
