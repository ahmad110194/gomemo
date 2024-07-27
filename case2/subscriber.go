package main

import (
	"log"
)

type Subscriber struct {
	event *Event
}

func NewSubscriber(event *Event) *Subscriber {
	return &Subscriber{event: event}
}

func (s *Subscriber) Start() {
	for {
		s.event.mu.Lock()
		for ch := range s.event.subscribers {
			select {
			case event := <-ch:
				log.Println("Received event:", event)
			}
		}
		s.event.mu.Unlock()
	}
}
