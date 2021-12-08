package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func Day8() {
	file, err := os.OpenFile("../input/day8.txt", os.O_RDONLY, os.ModePerm)

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

	log.Print(calculateOutput(entries))
}

func calculateOutput(entries []string) int {
	count := 0
	i := 0

	for i < len(entries) {
		entry := entries[i]
		components := strings.Split(entry, " | ")
		patterns := strings.Split(components[0], " ")
		output := strings.Split(components[1], " ")
		number := getDigits(patterns, output)
		count += number
		i += 1
	}

	return count
}

func getDigits(patterns []string, output []string) int {
	mapping := []map[string]bool{}
	sort.Slice(patterns, func(a, b int) bool {
		return len(patterns[a]) < len(patterns[b])
	})

	i := 0

	for i < len(patterns) {
		pattern := patterns[i]
		chars := strings.Split(pattern, "")
		j := 0
		mapp := make(map[string]bool)

		for j < len(chars) {
			char := chars[j]
			mapp[char] = true
			j += 1
		}

		mapping = append(mapping, mapp)
		i += 1
	}

	u := getUpper(mapping)

	midAndBottom := []string{}
	for key, _ := range mapping[3] {
		if key == u {
			continue
		}

		_, exists := mapping[4][key]
		_, exists2 := mapping[5][key]

		if exists && exists2 {
			midAndBottom = append(midAndBottom, key)
		}
	}

	_, exists := mapping[2][midAndBottom[0]]

	var mid string
	var bottom string

	if exists {
		mid = midAndBottom[0]
		bottom = midAndBottom[1]
	} else {
		mid = midAndBottom[1]
		bottom = midAndBottom[0]
	}

	var topLeft string
	for key, _ := range mapping[2] {
		if key == mid {
			continue
		}

		_, exists := mapping[0][key]

		if exists {
			continue
		} else {
			topLeft = key
			break
		}
	}

	i = 3
	for i < 6 {
		m := mapping[i]
		_, exists := m[topLeft]

		if exists {
			break
		}

		i += 1
	}

	var bottomRight string
	for key, _ := range mapping[i] {
		_, exists = mapping[0][key]
		if key == u || key == topLeft || key == bottom || key == mid {
			continue
		} else {
			bottomRight = key
			break
		}
	}

	var topRight string
	for key, _ := range mapping[0] {
		if key == bottomRight {
			continue
		} else {
			topRight = key
			break
		}
	}

	i = 3
	for i < 6 {
		m := mapping[i]

		_, exists := m[u]
		_, exist2 := m[topRight]
		_, exist3 := m[mid]
		_, exist4 := m[bottomRight]
		_, exist5 := m[bottom]

		if exists && exist2 && exist3 && !exist4 && exist5 {
			break
		}

		i += 1
	}

	var bottomLeft string
	for key, _ := range mapping[i] {
		if key == u || key == topRight || key == mid || key == bottom {
			continue
		} else {
			bottomLeft = key
			break
		}
	}

	log.Print("---")
	log.Printf("mid %v", mid)
	log.Printf("bottom %v", bottom)
	log.Printf("upper %v", u)
	log.Printf("topLeft %v", topLeft)
	log.Printf("bottom right %v", bottomRight)
	log.Printf("topRight %v", topRight)
	log.Printf("bottomLeft %v", bottomLeft)

	c := make(map[string]string)

	c[sort1(u+topRight+bottomRight+bottom+bottomLeft+topLeft)] = "0"
	c[sort1(topRight+bottomRight)] = "1"
	c[sort1(u+topRight+mid+bottomLeft+bottom)] = "2"
	c[sort1(u+topRight+mid+bottomRight+bottom)] = "3"
	c[sort1(topLeft+mid+topRight+bottomRight)] = "4"
	c[sort1(u+topLeft+mid+bottomRight+bottom)] = "5"
	c[sort1(u+topLeft+mid+bottomLeft+bottomRight+bottom)] = "6"
	c[sort1(u+topRight+bottomRight)] = "7"
	c[sort1("abcdefg")] = "8"
	c[sort1(u+topLeft+topRight+mid+bottomRight+bottom)] = "9"

	i = 0
	digits := []string{}

	for i < len(output) {
		o := output[i]
		digits = append(digits, c[sort1(o)])
		i += 1
	}

	number, err := strconv.Atoi(strings.Join(digits, ""))

	if err != nil {
		log.Fatal(err)
	}

	return number
}

func getUpper(mapping []map[string]bool) string {
	var u string

	for key, _ := range mapping[1] {
		_, exists := mapping[0][key]

		if !exists {
			u = key
			break
		}
	}

	return u
}

func sort1(s string) string {
	r := []rune(s)
	sort.Slice(r, func(i int, j int) bool { return r[i] < r[j] })
	return string(r)
}
