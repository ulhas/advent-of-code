package main

import (
	"log"
	"time"
)

func main() {
	start := time.Now()

	Day10()

	elapsed := time.Since(start)
	log.Printf("--- Time taken %s ---", elapsed)
}
