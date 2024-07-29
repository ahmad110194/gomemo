package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
)

func main() {
	go func() {
		fmt.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	app := App{}
	http.HandleFunc("/", app.Handler)
	fmt.Println("Starting server on :8080")
	http.ListenAndServe(":8080", nil)
}
