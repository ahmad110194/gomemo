package main

type Event struct {
	subscribers []chan string
}

func NewEvent() *Event {
	return &Event{}
}

func (eb *Event) Subscribe() chan string {
	ch := make(chan string)
	eb.subscribers = append(eb.subscribers, ch)
	return ch
}

func (eb *Event) Publish(event string) {
	for _, sub := range eb.subscribers {
		sub <- event
	}
}
