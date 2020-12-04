package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func day4() {
	file, err := os.OpenFile("./2020/input/day4.txt", os.O_RDONLY, os.ModePerm)

	if err != nil {
		log.Fatalf("Open file error: %v", err)
		return
	}

	defer file.Close()

	reader := bufio.NewReader(file)
	passports := []string{}
	passport := ""

	for {
		line, err := reader.ReadString('\n')

		if err != nil {
			if err == io.EOF {
				trimmedPassport := strings.TrimSpace(passport)
				passports = append(passports, trimmedPassport)
				break
			}

			log.Fatalf("Read file line error: %v", err)
			return
		}

		trimmed := strings.TrimSpace(line)
		log.Printf("Input -%v-", trimmed)

		if trimmed == "" {
			trimmedPassport := strings.TrimSpace(passport)
			passports = append(passports, trimmedPassport)
			passport = ""
		} else {
			passport += " " + trimmed
		}

	}

	count := 0

	for i := 0; i < len(passports); i++ {
		if isValidPassport(passports[i]) {
			count++
		}
	}

	log.Printf("Count %v", count)
}

func isValidPassport(passport string) bool {
	log.Printf("Passport %v", passport)
	components := strings.Split(passport, " ")

	byr := false
	iyr := false
	eyr := false
	hgt := false
	hcl := false
	ecl := false
	pid := false

	for i := 0; i < len(components); i++ {
		comps := strings.Split(components[i], ":")
		comp := comps[0]
		value := comps[1]

		if comp == "byr" {
			number, convError := strconv.Atoi(value)

			if convError != nil {
				log.Fatalf("Conversion error: %v", convError)
				byr = false
			}

			byr = number >= 1920 && number <= 2002
			log.Printf("Byr %v %v", number, byr)
		} else if comp == "iyr" {
			number, convError := strconv.Atoi(value)

			if convError != nil {
				log.Fatalf("Conversion error: %v", convError)
				iyr = false
			}

			iyr = number >= 2010 && number <= 2020
			log.Printf("IYR %v %v", number, iyr)
		} else if comp == "eyr" {
			number, convError := strconv.Atoi(value)

			if convError != nil {
				log.Fatalf("Conversion error: %v", convError)
				eyr = false
			}

			eyr = number >= 2020 && number <= 2030
			log.Printf("Eyr %v %v", number, eyr)
		} else if comp == "hgt" {
			if strings.HasSuffix(value, "in") {
				inches := strings.Split(value, "in")[0]
				number, convError := strconv.Atoi(inches)

				if convError != nil {
					log.Fatalf("Conversion error: %v", convError)
					hgt = false
				}

				hgt = number >= 59 && number <= 76
				log.Printf("Inch Number %v %v", number, hgt)
			} else if strings.HasSuffix(value, "cm") {
				cms := strings.Split(value, "cm")[0]
				number, convError := strconv.Atoi(cms)

				if convError != nil {
					log.Fatalf("Conversion error: %v", convError)
					hgt = false
				}

				hgt = number >= 150 && number <= 193
				log.Printf("CM Number %v %v", number, hgt)
			} else {
				log.Printf("No inches and cm")
				hgt = false
			}
		} else if comp == "hcl" {
			if !(strings.HasPrefix(value, "#")) {
				hcl = false
			} else {
				hclValue := strings.Split(value, "#")[1]

				if len(hclValue) != 6 {
					log.Printf("HCL Value %v %v", hclValue, false)
					hcl = false
				} else {
					characterSet := "0123456789abcdef"
					c := strings.Split(hclValue, "")
					hcl = true

					for i := 0; i < len(c); i++ {
						if !strings.Contains(characterSet, c[i]) {
							hcl = false
							break
						}
					}

					log.Printf("HCL Value %v %v", hclValue, hcl)
				}
			}
		} else if comp == "ecl" {
			ecl = value == "amb" || value == "blu" || value == "brn" || value == "gry" || value == "grn" || value == "hzl" || value == "oth"
			log.Printf("ECL %v %v", value, ecl)
		} else if comp == "pid" {
			pid = len(value) == 9
			log.Printf("PID %v %v", value, pid)
		}
	}

	result := byr && iyr && eyr && hgt && ecl && hcl && pid
	log.Printf("Valid %v", result)
	return result
}
