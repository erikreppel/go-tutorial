package main

import (
	"log"
	"time"

	"github.com/erikreppel/go-tutorial/5_concurrency/onetomany/example"
)

func main() {
	start := time.Now()
	example.SerialDouble(100, false)
	log.Println("Serial took:", time.Since(start))

	start = time.Now()
	example.ConcurrentDouble(100, 10, false)
	log.Println("Concurrent took:", time.Since(start))
}
