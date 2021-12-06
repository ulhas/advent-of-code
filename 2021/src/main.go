package main

import (
	"log"
	"time"
)

func main() {
	start := time.Now()

	Day6()

	elapsed := time.Since(start)
	log.Printf("--- Time taken %s ---", elapsed)
}
