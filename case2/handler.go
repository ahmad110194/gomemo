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
	key := r.URL.Query().Get("key")

	ch := a.event.Subscribe()
	defer a.event.Unsubscribe(ch)

	a.publishEvent(key)
}

func (a *App) publishEvent(message string) {
	a.event.Publish("New event: " + message)
}
