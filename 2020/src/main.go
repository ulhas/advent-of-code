package main

import (
	"log"
	"time"
)

func main() {
	start := time.Now()

	day6()

	elapsed := time.Since(start)
	log.Printf("--- Time taken %s ---", elapsed)
}
