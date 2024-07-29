package main

import "sync"

type Event struct {
	subscribers map[chan string]struct{}
	mu          sync.Mutex
}

func NewEvent() *Event {
	return &Event{
		subscribers: make(map[chan string]struct{}),
	}
}

func (eb *Event) Subscribe() chan string {
	ch := make(chan string)
	eb.mu.Lock()
	eb.subscribers[ch] = struct{}{}
	eb.mu.Unlock()
	return ch
}

func (eb *Event) Unsubscribe(ch chan string) {
	eb.mu.Lock()
	defer eb.mu.Unlock()
	if _, ok := eb.subscribers[ch]; ok {
		delete(eb.subscribers, ch)
	}
}

func (eb *Event) Publish(event string) {
	eb.mu.Lock()
	defer eb.mu.Unlock()
	for sub := range eb.subscribers {
		sub <- event
	}
}
