package main

import (
	"sync"
)

type Event struct {
	subscribers map[int]*Subscriber
	mu          sync.Mutex
	nextID      int
}

func NewEvent() *Event {
	return &Event{
		subscribers: make(map[int]*Subscriber),
		nextID:      0,
	}
}

func (e *Event) Subscribe() *Subscriber {
	e.mu.Lock()
	defer e.mu.Unlock()

	sub := NewSubscriber(e.nextID)
	e.subscribers[e.nextID] = sub
	e.nextID++

	return sub
}

func (e *Event) Unsubscribe(sub *Subscriber) {
	e.mu.Lock()
	defer e.mu.Unlock()

	delete(e.subscribers, sub.id)
}

func (e *Event) Publish(message string) {
	e.mu.Lock()
	defer e.mu.Unlock()

	for _, sub := range e.subscribers {
		sub.channel <- message
	}
}
