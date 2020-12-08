package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

// Instruction is instruction
type Instruction struct {
	text       string
	value      int
	hasRunOnce bool
}

func day8() {
	file, err := os.OpenFile("./2020/input/day8.txt", os.O_RDONLY, os.ModePerm)

	if err != nil {
		log.Fatalf("Open file error: %v", err)
		return
	}

	defer file.Close()

	reader := bufio.NewReader(file)
	instructions := []Instruction{}

	for {
		line, err := reader.ReadString('\n')

		if err != nil {
			if err == io.EOF {
				break
			}

			log.Fatalf("Read file line error: %v", err)
			return
		}

		input := strings.TrimSpace(line)
		components := strings.Split(input, " ")

		value, convError := strconv.Atoi(components[1])

		if convError != nil {
			log.Fatalf("Conversion error: %v", convError)
			return
		}

		instruction := Instruction{text: components[0], value: value, hasRunOnce: false}
		instructions = append(instructions, instruction)
	}

	log.Printf("Accumulator %v", getAccumulatorValue(instructions))
}

func getAccumulatorValue(instructions []Instruction) int {
	accumulator := 0
	i := 0

	for true {
		instruction := &instructions[i]

		if instruction.hasRunOnce {
			break
		}

		if instruction.text == "jmp" {
			i += instruction.value
		} else if instruction.text == "acc" {
			accumulator += instruction.value
			i++
		} else {
			i++
		}

		instruction.hasRunOnce = true
	}

	return accumulator
}
