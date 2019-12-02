package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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

	trimmedLine := strings.TrimSpace(line)

	fmt.Println("********** Day 2 **********")
	fmt.Printf("Trimmed Line :%v \n", trimmedLine)
	fmt.Println("********** End Day 2 **********")
}
