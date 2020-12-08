package main

import (
	"log"
	"time"
)

func main() {
	start := time.Now()

	day8()

	elapsed := time.Since(start)
	log.Printf("--- Time taken %s ---", elapsed)
}
