package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	_ "net/http/pprof"

	"github.com/newrelic/go-agent/v3/newrelic"
)

func main() {
	newrelicApp, err := newrelic.NewApplication(
		newrelic.ConfigAppName("local-testing-case-1"),
		newrelic.ConfigLicense("__YOUR_NEW_RELIC_LICENSE_KEY__1234567890"),
	)
	if err != nil {
		fmt.Println("Error initializing New Relic:", err)
		return
	}

	app := &App{newrelicApp: newrelicApp}

	go func() {
		fmt.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	http.HandleFunc("/", app.handler)
	fmt.Println("Starting server on :8080")
	http.ListenAndServe(":8080", nil)
}

type App struct {
	newrelicApp *newrelic.Application
}

func (a *App) handler(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithCancel(r.Context())
	defer cancel()

	newrelicTx := a.newrelicApp.StartTransaction("transaction_name")
	defer newrelicTx.End()

	ctxWithNewRelicTransaction := newrelic.NewContext(ctx, newrelicTx)

	go processRequest(ctxWithNewRelicTransaction)

	// Simulate a delay to keep the handler active longer
	time.Sleep(10 * time.Second)
}

func processRequest(ctx context.Context) {
	ctxReq, _ := context.WithCancel(ctx)
	doSomeProcess(ctxReq)
}

func doSomeProcess(ctx context.Context) {
	// simulate some process
	for {
		select {
		case <-ctx.Done():
			return
		default:
			time.Sleep(3 * time.Second) // simulate processing
		}
	}
}
