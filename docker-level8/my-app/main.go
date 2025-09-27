package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"time"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, from the docker world!")
}

func slogHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, from the docker world with slog!")
	slog.Info("slog handler was called")
}

func eventsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")
	w.Header().Set("X-Accel-Buffering", "no")

	flusher, ok := w.(http.Flusher)
	if !ok {
		http.Error(w, "Streaming unsupported!", http.StatusInternalServerError)
		return
	}

	fmt.Fprint(w, ": hello\n\n")
	fmt.Fprint(w, "retry: 3000\n")
	flusher.Flush()

	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	ctx := r.Context()

	for {
		select {
		case t := <-ticker.C:
			fmt.Fprintf(w, "data: %s\n\n", t.Format(time.RFC3339))
		case <-ctx.Done():
			return
		}
	}
}

func main() {
	http.HandleFunc("/", helloHandler)
	http.HandleFunc("/slog", slogHandler)
	http.HandleFunc("/events", eventsHandler)
	slog.Info("Starting server on :8080")
	slog.Info("listening on http://localhost:8080/events for SSE")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		slog.Error("Error starting server:", "error", err)
	}
}
