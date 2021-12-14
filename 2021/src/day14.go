package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"strings"
)

type Polymer struct {
	template  map[string]uint64
	histogram map[string]uint64
}

func Day14() {
	file, err := os.OpenFile("../input/day14.txt", os.O_RDONLY, os.ModePerm)

	if err != nil {
		log.Fatalf("Open file error: %v", err)
		return
	}

	defer file.Close()

	reader := bufio.NewReader(file)
	var template string
	pairInsertions := make(map[string]string)
	i := 0

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

		if i == 0 {
			template = trimmed
		} else if trimmed != "" {
			components := strings.Split(trimmed, " -> ")
			pairInsertions[components[0]] = components[1]
		}

		i += 1
	}

	polymer := createPolymer(template)

	i = 0
	for i < 400 {
		polymer = polymer.step(pairInsertions)
		i += 1
	}

	log.Print(polymer.getDiff())
}

func createPolymer(template string) Polymer {
	frequency := make(map[string]uint64)
	histogram := make(map[string]uint64)
	i := 0

	for i < len(template)-1 {
		key := template[i : i+2]
		value, exists := frequency[key]

		if exists {
			frequency[key] = value + 1
		} else {
			frequency[key] = 1
		}

		key = string(template[i])
		value, exists = histogram[key]
		if exists {
			histogram[key] = value + 1
		} else {
			histogram[key] = 1
		}

		i += 1
	}

	key := string(template[len(template)-1])
	value, exists := histogram[key]
	if exists {
		histogram[key] = value + 1
	} else {
		histogram[key] = 1
	}

	return Polymer{template: frequency, histogram: histogram}
}

func (polymer Polymer) step(pairInsertions map[string]string) Polymer {
	histogram := polymer.histogram
	template := make(map[string]uint64)

	for pair, insertion := range pairInsertions {
		count, exists := polymer.template[pair]

		if exists {
			components := strings.Split(pair, "")
			pair1 := fmt.Sprintf("%s%s", components[0], insertion)
			pair2 := fmt.Sprintf("%s%s", insertion, components[1])

			freq1, exists := template[pair1]
			if exists {
				template[pair1] = freq1 + count
			} else {
				template[pair1] = count
			}

			freq2, exists := template[pair2]
			if exists {
				template[pair2] = freq2 + count
			} else {
				template[pair2] = count
			}

			value, exists := histogram[insertion]
			if exists {
				histogram[insertion] = value + count
			} else {
				histogram[insertion] = count
			}
		}
	}

	return Polymer{template: template, histogram: histogram}
}

func (polymer Polymer) getDiff() uint64 {
	var max uint64 = 0
	var min uint64 = math.MaxUint64

	for _, value := range polymer.histogram {
		if value > max {
			max = value
		}

		if value < min {
			min = value
		}
	}

	return max - min
}

// func getFrequency(polymer Polymer) Polymer {

// 	for key, value := range pairInsertions {
// 		freq, exists := frequency[key]

// 		if exists {
// 			components := strings.Split(key, "")
// 			key1 := fmt.Sprintf("%s%s", components[0], value)
// 			key2 := fmt.Sprintf("%s%s", value, components[1])

// 			freq1, exists := newFrequency[key1]
// 			if exists {
// 				newFrequency[key1] = freq1 + freq
// 			} else {
// 				newFrequency[key1] = freq
// 			}

// 			freq2, exists := newFrequency[key2]
// 			if exists {
// 				newFrequency[key2] = freq2 + freq
// 			} else {
// 				newFrequency[key2] = freq
// 			}

// 			his, exists := histogram[value]
// 			if exists {
// 				histogram[value] = freq + his
// 			} else {
// 				histogram[value] = freq
// 			}
// 		}
// 	}

// 	log.Printf("Histogram %v", histogram)
// 	log.Printf("New Frequency %v", newFrequency)

// 	return newFrequency
// }

// func getHistogram(frequency map[string]uint64) map[string]uint64 {
// 	newFrequency := make(map[string]uint64)

// 	for key, value := range frequency {
// 		components := strings.Split(key, "")

// 		key1 := components[0]
// 		freq1, exists := newFrequency[key1]

// 		if exists {
// 			newFrequency[key1] = freq1 + value
// 		} else {
// 			newFrequency[key1] = value
// 		}

// 		key2 := components[1]
// 		freq2, exists := newFrequency[key2]

// 		if exists {
// 			newFrequency[key2] = freq2 + value
// 		} else {
// 			newFrequency[key2] = value
// 		}
// 	}

// 	return newFrequency
// }
