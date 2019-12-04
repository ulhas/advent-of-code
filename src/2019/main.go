package main

import (
	"log"
	"time"
)

func main() {
	start := time.Now()

	day1()
	day2()
	day3()

	elapsed := time.Since(start)
	log.Printf("--- Time taken %s ---", elapsed)
}
