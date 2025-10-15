# **Task 6 – Dockerize Your Go App 🐹**

**Goal**: Write a tiny Go web server and run it inside Docker.

Steps:

1. Write a Go program (`main.go`) that runs an HTTP server on port `8080`. (Simple “Hello from uwubuntu!” is enough.)
2. Create a Dockerfile:

   * `FROM golang:1.23` (or latest Go image).
   * `WORKDIR /app`
   * `COPY . .`
   * `RUN go build -o server .`
   * `CMD ["./server"]`
3. Build it: `docker build -t my-go-server .`
4. Run it: `docker run -p 8080:8080 my-go-server`.
5. Open `http://localhost:8080` and see your Go app running from inside Docker.

💡 Hints:

* Your Go app must `ListenAndServe(":8080")`.
* If something doesn’t work, check `docker logs <container_id>`.

