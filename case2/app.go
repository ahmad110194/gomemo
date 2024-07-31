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

	sub := a.event.Subscribe()

	go sub.Start()

	a.publishEvent(key)

	a.event.Unsubscribe(sub)
}

func (a *App) publishEvent(message string) {
	a.event.Publish("New event: " + message)
}
