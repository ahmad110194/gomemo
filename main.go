package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"sync"
	"time"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func main() {
	rand.Seed(time.Now().UnixNano())

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

	str := randString(10)
	resp, err := http.Get(fmt.Sprintf("http://localhost:8080/?key=key_%s&value=value_%s", str, str))
	if err != nil {
		fmt.Printf("Error making request: %v\n", err)
		return
	}
	defer resp.Body.Close()
}

func randString(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}
