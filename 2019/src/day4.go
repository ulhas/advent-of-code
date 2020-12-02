package main

import "log"

func day4() {
	inputStart := 248345
	inputEnd := 746315
	numberOfValidPasswords := 0

	for i := inputStart; i <= inputEnd; i++ {
		if meets(i) {
			numberOfValidPasswords += 1
		}
	}

	log.Printf("Number of valid password %v", numberOfValidPasswords)
}

func meets(number int) bool {
	ones, tens, hundreds, thousands, tenThousands, hundredThousands := digits(number)

	if ones == tens || tens == hundreds || hundreds == thousands || thousands == tenThousands || tenThousands == hundredThousands {
		return hundredThousands <= tenThousands && tenThousands <= thousands && thousands <= hundreds && hundreds <= tens && tens <= ones
	}

	return false
}

func digits(number int) (ones int, tens int, hundreds int, thousands int, tenThousands int, hundredThousands int) {
	ones = number % 10
	tens = (number / 10) % 10
	hundreds = (number / 100) % 10
	thousands = (number / 1000) % 10
	tenThousands = (number / 10000) % 10
	hundredThousands = (number / 100000) % 10

	return ones, tens, hundreds, thousands, tenThousands, hundredThousands
}
