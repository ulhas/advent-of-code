package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"sort"
	"strings"
)

func Day10() {
	file, err := os.OpenFile("../input/day10.txt", os.O_RDONLY, os.ModePerm)

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

	// log.Print(getErrorScore(entries))
	log.Print(getMiddleScore(entries))
}

func getErrorScore(entries []string) int {
	corresponding := map[string]string{"}": "{", ")": "(", ">": "<", "]": "["}
	score := 0
	i := 0
	corrupted := []string{}

	for i < len(entries) {
		entry := entries[i]
		stack := []string{}
		components := strings.Split(entry, "")
		j := 0

		for j < len(components) {
			ch := components[j]

			if ch == "{" || ch == "<" || ch == "(" || ch == "[" {
				stack = append(stack, ch)
			} else {
				n := len(stack) - 1
				c := stack[n]
				value, _ := corresponding[ch]

				if c != value {
					corrupted = append(corrupted, ch)
					break
				} else {
					stack = stack[:n]
				}
			}

			j += 1
		}

		i += 1
	}

	scores := map[string]int{")": 3, "]": 57, "}": 1197, ">": 25137}
	i = 0

	for i < len(corrupted) {
		ch := corrupted[i]
		value, _ := scores[ch]
		score += value
		i += 1
	}

	return score
}

func getMiddleScore(entries []string) int {
	corresponding := map[string]string{"}": "{", ")": "(", ">": "<", "]": "["}
	i := 0
	stacks := [][]string{}

	for i < len(entries) {
		entry := entries[i]
		stack := []string{}
		components := strings.Split(entry, "")
		j := 0
		corrupted := false

		for j < len(components) {
			ch := components[j]

			if ch == "{" || ch == "<" || ch == "(" || ch == "[" {
				stack = append(stack, ch)
			} else {
				n := len(stack) - 1
				c := stack[n]
				value, _ := corresponding[ch]

				if c != value {
					corrupted = true
					break
				} else {
					stack = stack[:n]
				}
			}

			j += 1
		}

		if !corrupted {
			stacks = append(stacks, stack)
		}

		i += 1
	}

	points := map[string]int{")": 1, "]": 2, "}": 3, ">": 4}
	corresponding = map[string]string{"{": "}", "(": ")", "<": ">", "[": "]"}
	scores := []int{}

	for _, row := range stacks {
		score := 0
		j := len(row) - 1

		for j >= 0 {
			cell := row[j]
			ch, _ := corresponding[cell]
			score = score*5 + points[ch]
			j -= 1
		}

		scores = append(scores, score)
	}

	sort.Ints(scores)

	return scores[len(scores)/2]
}
