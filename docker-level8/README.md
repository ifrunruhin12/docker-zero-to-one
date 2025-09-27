# Level 8: The Haunted House of Ports

## Task 8 overview and experiments

So, I have been doing some experiments with ports and there's some stuff I came across which seemed a bit confusing.
First of all, I created a directory for this level and then created a my-app directory inside it.
I created a simple web app which has three handlers. One is in the home page it will say hello from docker, one will show the same thing from slog in /slog api and another
is in the /events which will stream SSE events and show current time and update it every seconds.
Here's the code I written:

```go
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

```

So, first I ran it using `docker run -p 9090:9090 my-app` outside the my-app directory. But it said Unable to find image 'my-app:latest' locally.
and then give me an error saying requested access to the resource is denied. then I understod that I need to create a Dockerfile to build the app container first?
So, I created this docker file below:

```Dockerfile
FROM golang:1.25.1
WORKDIR /app
COPY . .
RUN go build -o server .
CMD ["./server"]

```

I ran it using `docker buildx build -t cool-web-app .` but then it gave me error saying there's no `go.mod` file found
in the current directory, well, because the Dockerfile was outside of my `my-app` directory. So I learned Dockerfile should be in the same
directory is my `main.go` and `go.mod` file. Anyway so I moved it inside the `my-app` folder and ran `docker buildx build -t cool-web-app .` again
and this time it worked and built it. then after the build was complete I ran it using `docker run -p 9090:9090 cool-web-app` and it said it was running giving me messages like this

```
2025/09/24 05:47:10 INFO starting server on :8080
2025/09024 05:47:10 INFO listening on http://localhost:8080/events for SSE
```

but I didn't see anything when I went to 9090 on my host. so I then understood so maybe the docker container port and the port my go app is listening to
have to be the same? So when I ran it again with `docker run -p 9090:8080 cool-web-app` and now it worked and I can see the messages when I went to localhost:9090
on my host. but I noticed something odd when I went to /events api but for some reason it wasn't loading. So I tried it with `go run main.go` but on host `8080` now the /events
was still loading. so I run it with docker again and tested the `9090` in port in postman. But although the /events was giving 200 ok status it wasn't updating the time every seconds
or even if it did updated the time every seconds I wasn't able to see it.

Anyway then I tried, running nginx in `9090` port on the host where my go app was already running so it gave me error saying the port is alrready in use.
So, I learned you cannot use the same port for host. but same port in the container for the same service `nginx` is fine?
so then I run two nginx first ran `docker run -p 8080:80 nginx` and then opened localhost on port 8080 and yes I could have seen the nginx welcome page.
then again did `docker run -p 8081:80 nginx` and then opened localhost on port 8081 and yes I can see the welcome page of nginx. but after that I ran
`docker run -p 8082:8080 nginx` and from terminal log it seemed like nginx running just fine but when I opened 8082 on my host I saw nothing and it said connection was reset.
why tho? Does nginx doesn't run without the port `:80`?
