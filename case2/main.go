package main

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
)

func main() {
	event := NewEvent()
	subscriber := NewSubscriber(event)
	app := NewApp(event)

	go subscriber.Start()

	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	http.HandleFunc("/", app.Handler)
	fmt.Println("Starting server on :8080")
	http.ListenAndServe(":8080", nil)
}
