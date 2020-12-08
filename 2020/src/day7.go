package main

import (
	"bufio"
	"errors"
	"io"
	"log"
	"os"
	"strings"
)

// Bag is a bag and chldren
type Bag struct {
	color             string
	hasShinyGoldChild bool
	children          []Bag
}

func day7() {
	file, err := os.OpenFile("./2020/input/day7.txt", os.O_RDONLY, os.ModePerm)

	if err != nil {
		log.Fatalf("Open file error: %v", err)
		return
	}

	defer file.Close()

	reader := bufio.NewReader(file)
	global := Bag{color: "start", hasShinyGoldChild: false, children: []Bag{}}

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
		bag := getBag(input)
		traversed, err := traverse(global, bag.color)

		if err != nil {
			// log.Printf("No bag found. Appending to global.")
			global.children = append(global.children, bag)
		} else {
			log.Printf("Bag Found. Appending to child")
			traversed.children = append(traversed.children, bag)
		}
	}

	log.Printf("Bag Children length %v", len(global.children))
}

func getBag(input string) Bag {
	components := strings.Split(input, " contain ")

	outerBag := components[0]
	outerBagSplit := strings.Split(outerBag, " ")
	outerBagColor := outerBagSplit[0] + " " + outerBagSplit[1]

	bag := Bag{color: outerBagColor, hasShinyGoldChild: false, children: []Bag{}}

	innerBag := components[1]

	if innerBag == "no other bags." {
		return bag
	}

	untrimmedInnerBags := strings.Split(innerBag, ", ")

	children := []Bag{}
	for i := 0; i < len(untrimmedInnerBags); i++ {
		innerBag := strings.TrimSuffix(untrimmedInnerBags[i], ".")

		innerBagSplit := strings.Split(innerBag, " ")
		innerBagColor := innerBagSplit[1] + " " + innerBagSplit[2]

		if innerBagColor == "shiny gold" {
			bag.hasShinyGoldChild = true
		}

		children = append(children, Bag{color: innerBagColor, children: []Bag{}})
	}
	bag.children = children

	return bag
}

func traverse(bag Bag, color string) (*Bag, error) {
	children := bag.children

	for i := 0; i < len(children); i++ {
		b := children[i]

		if b.color == color {
			return &b, nil
		}

		return traverse(b, color)
	}

	return &bag, errors.New("bag_not_found")
}
