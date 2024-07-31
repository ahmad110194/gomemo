package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"sync"
	"time"
)

type App struct {
	attemp int
	mu     sync.Mutex
	tasks  map[string]*Task
}

type Task struct {
	id     string
	status string
}

func (a *App) taskRetryer(task *Task) {
	duration := time.Duration(rand.Intn(10)) * time.Second
	time.Sleep(duration)

	if duration >= 1*time.Second {
		a.mu.Lock()
		a.attemp += 1
		a.mu.Unlock()

		go a.taskRetryer(task)
	} else {
		a.mu.Lock()
		task.status = "done"
		a.mu.Unlock()

		fmt.Printf("Task completed at attemp %d\n", a.attemp)
	}
}

func (a *App) Handler(w http.ResponseWriter, r *http.Request) {
	task := &Task{
		id:     fmt.Sprintf("task-%d", time.Now().UnixNano()),
		status: "running",
	}
	go a.taskRetryer(task)
}
