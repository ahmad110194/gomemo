package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	for {
		wg.Add(1)
		go makeRequest(&wg)

		delay := time.Duration(rand.Intn(100)) * time.Millisecond
		time.Sleep(delay)
	}
	wg.Wait()
}

func makeRequest(wg *sync.WaitGroup) {
	defer wg.Done()
	resp, err := http.Get("http://localhost:8080")
	if err != nil {
		fmt.Printf("Error making request: %v\n", err)
		return
	}
	defer resp.Body.Close()
}
