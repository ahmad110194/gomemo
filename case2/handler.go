package main

import (
	"net/http"
)

type App struct {
	event *Event
}

func NewApp(event *Event) *App {
	return &App{event: event}
}

func (a *App) Handler(w http.ResponseWriter, r *http.Request) {
	a.publishEvent()
	// fmt.Fprintln(w, "Event published")
}

func (a *App) publishEvent() {
	a.event.Publish("New event")
}
