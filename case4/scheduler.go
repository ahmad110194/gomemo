package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"sync"
	"time"
)

type Scheduler struct {
}

func (s *Scheduler) ScheduleTask(depth int, wg *sync.WaitGroup) {
	defer wg.Done()
	if depth > 0 {
		wg.Add(1)
		go func(d int) {
			sleepDuration := time.Duration(rand.Intn(5)) * time.Second
			time.Sleep(sleepDuration)
			fmt.Printf("Task completed at depth %d, sleep duration %v\n", d, sleepDuration)
			s.ScheduleTask(d-1, wg)
		}(depth)
	}
}

// Handler HTTP untuk menjadwalkan tugas
func ScheduleHandler(w http.ResponseWriter, r *http.Request) {
	scheduler := &Scheduler{}
	var wg sync.WaitGroup
	wg.Add(1)
	go scheduler.ScheduleTask(10, &wg)
	wg.Wait()
	fmt.Fprintf(w, "Tugas dijadwalkan")
}
