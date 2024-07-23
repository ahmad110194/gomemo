package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
)

type App struct {
	c *Cache
}

func (a *App) handler(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Query().Get("key")
	value := r.URL.Query().Get("value")
	a.c.Set(key, value)

	val, _ := a.c.Get(key)
	fmt.Fprintf(w, "Cache %s = %s\n", key, val)
}

func NewApp(c *Cache) *App {
	return &App{
		c: c,
	}
}

func main() {
	c := NewCache()
	app := NewApp(c)

	go func() {
		fmt.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	http.HandleFunc("/", app.handler)
	fmt.Println("Starting server on :8080")
	http.ListenAndServe(":8080", nil)
}
