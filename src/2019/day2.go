package main

import (
	"bufio"
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
	log.Printf("Input IntCode :%v", input)
	found := false

	for i := 0; i <= 99; i++ {
		if found {
			break
		}

		for j := 0; j <= 99; j++ {
			temp := make([]int, len(input))
			copy(temp, input)

			temp[1] = i
			temp[2] = j

			output := getCalculatedIntCode(temp)

			if output[0] == 19690720 {
				log.Printf("Noun is :%v, Verb is :%v", i, j)
				found = true
				break
			}
		}
	}
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
			log.Print("\n\nOpcode 99 encountered. Aborting")
			return intCode
		}

		if opCode == 1 || opCode == 2 {
			firstParameter := intCode[position+1]
			secondParameter := intCode[position+2]
			thirdParameter := intCode[position+3]

			if opCode == 1 {
				intCode[thirdParameter] = intCode[firstParameter] + intCode[secondParameter]
			} else {
				intCode[thirdParameter] = intCode[firstParameter] * intCode[secondParameter]
			}
		} else {
			log.Printf("\n\nInvalid OpCode :%v", opCode)
			break
		}

		position += 4
	}

	return intCode
}
