package main

import (
	"log"
	"time"
)

func main() {
	start := time.Now()

	// day1()
	// day2()
	day3()

	elapsed := time.Since(start)
	log.Printf("\n\n\n --- Time taken %s ---", elapsed)
}
