package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"strings"
)

func day6() {
	file, err := os.OpenFile("./2020/input/day6.txt", os.O_RDONLY, os.ModePerm)

	if err != nil {
		log.Fatalf("Open file error: %v", err)
		return
	}

	defer file.Close()

	reader := bufio.NewReader(file)
	groups := []string{}
	group := ""

	for {
		line, err := reader.ReadString('\n')

		if err != nil {
			if err == io.EOF {
				trimmedGroup := strings.TrimSpace(group)
				groups = append(groups, trimmedGroup)
				break
			}

			log.Fatalf("Read file line error: %v", err)
			return
		}

		input := strings.TrimSpace(line)

		if input == "" {
			trimmedGroup := strings.TrimSpace(group)
			groups = append(groups, trimmedGroup)
			group = ""
			continue
		} else {
			group += " " + input
		}
	}

	count := 0

	for i := 0; i < len(groups); i++ {
		subGroups := groups[i]

		if len(subGroups) == 1 {
			count++
		} else {
			components := strings.Split(subGroups, " ")
			componentLength := len(components)
			firstSubGroup := strings.Split(components[0], "")

			for j := 0; j < len(firstSubGroup); j++ {
				char := firstSubGroup[j]
				present := true

				for k := 1; k < componentLength; k++ {
					if !strings.Contains(components[k], char) {
						present = false
						break
					}
				}

				if present {
					count++
				}
			}
		}
	}

	log.Printf("Count %v", count)
}
