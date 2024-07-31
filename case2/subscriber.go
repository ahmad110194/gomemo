package main

import (
	"fmt"
	"time"
)

type Subscriber struct {
	id       int
	channel  chan string
	stopChan chan struct{}
}

func NewSubscriber(id int) *Subscriber {
	return &Subscriber{
		id:       id,
		channel:  make(chan string, 10),
		stopChan: make(chan struct{}),
	}
}

func (s *Subscriber) Start() {
	for {
		select {
		case message := <-s.channel:
			time.Sleep(100 * time.Millisecond)
			fmt.Printf("Received message: %s, subid: %d\n", message, s.id)
		case <-s.stopChan:
			return
		}
	}
}
