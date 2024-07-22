package main

import (
	"flag"
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"sync"
	"time"
)

func main() {
	iterations := flag.Int("i", 1, "Number of iterations to make HTTP requests")
	flag.Parse()

	if *iterations < 1 {
		fmt.Println("Iterations must be greater than 0")
		os.Exit(1)
	}

	var wg sync.WaitGroup
	for i := 0; i < *iterations; i++ {
		wg.Add(1)

		delay := time.Duration(rand.Intn(300)) * time.Millisecond
		time.Sleep(delay)
		go makeRequest(i, &wg)
	}

	wg.Wait()
}

func makeRequest(index int, wg *sync.WaitGroup) {
	defer wg.Done()

	resp, err := http.Get("http://localhost:8080")
	if err != nil {
		fmt.Printf("Error making request %d: %v\n", index+1, err)
		return
	}
	defer resp.Body.Close()
}
