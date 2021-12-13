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

type Fold struct {
	direction string
	at        int
}

func Day13() {
	file, err := os.OpenFile("../input/day13.txt", os.O_RDONLY, os.ModePerm)

	if err != nil {
		log.Fatalf("Open file error: %v", err)
		return
	}

	defer file.Close()

	reader := bufio.NewReader(file)
	dots := make(map[string]bool)
	folds := []Fold{}

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

		if trimmed != "" {
			if strings.HasPrefix(trimmed, "fold") {
				components := strings.Split(trimmed, " ")
				components = strings.Split(components[2], "=")

				value, err := strconv.Atoi(components[1])

				if err != nil {
					log.Fatal("Not a number")
				}

				folds = append(folds, Fold{direction: components[0], at: value})
			} else {
				dots[trimmed] = true
			}
		}
	}

	i := 0

	for i < len(folds) {
		fold := folds[i]
		dots = fold.execute(dots)
		i += 1
	}

	i = 0

	for i < 40 {
		j := 0

		for j < 10 {
			key := fmt.Sprintf("%d,%d", i, j)

			_, exists := dots[key]

			if !exists {
				fmt.Print(".  ")
			} else {
				fmt.Print("#  ")
			}

			j += 1
		}

		fmt.Print("\n")
		i += 1
	}
}

func (fold Fold) up(dots map[string]bool) map[string]bool {
	newDots := make(map[string]bool)

	for key, _ := range dots {
		components := strings.Split(key, ",")

		x, err := strconv.Atoi(components[0])
		if err != nil {
			log.Fatal("x cant be converted")
		}

		y, err := strconv.Atoi(components[1])
		if err != nil {
			log.Fatal("x cant be converted")
		}

		if y < fold.at {
			newDots[key] = true
			continue
		}

		y = 2*fold.at - y
		newDots[fmt.Sprintf("%d,%d", x, y)] = true
	}

	return newDots
}

func (fold Fold) left(dots map[string]bool) map[string]bool {
	newDots := make(map[string]bool)

	for key, _ := range dots {
		components := strings.Split(key, ",")

		x, err := strconv.Atoi(components[0])
		if err != nil {
			log.Fatal("x cant be converted")
		}

		y, err := strconv.Atoi(components[1])
		if err != nil {
			log.Fatal("x cant be converted")
		}

		if x < fold.at {
			newDots[key] = true
			continue
		}

		x = 2*fold.at - x
		newDots[fmt.Sprintf("%d,%d", x, y)] = true
	}

	return newDots
}

func (fold Fold) execute(dots map[string]bool) map[string]bool {
	if fold.direction == "y" {
		return fold.up(dots)
	}

	return fold.left(dots)
}
