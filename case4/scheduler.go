package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

type App struct {
	attemp int
}

func (a *App) taskRetryer() {
	duration := time.Duration(rand.Intn(10)) * time.Second
	time.Sleep(duration)

	if duration >= 1*time.Second {
		a.attemp += 1
		a.taskRetryer()
	}

	fmt.Printf("Task completed at attemp %d\n", a.attemp)
}

func (a *App) Handler(w http.ResponseWriter, r *http.Request) {
	go a.taskRetryer()
}
