## 🧪 Task 11 – Slim Your Go App

**Goal**: Refactor your Level 10 Go API to use a multi-stage build.

### Steps:

1. Navigate to your `docker-level10/api/` directory.
2. Update the `Dockerfile` to use multi-stage builds:
   - **Stage 1**: Build the binary (use `golang:1.25.1`).
   - **Stage 2**: Run it (use `alpine:latest`).
3. Build the new image: `docker build -t api-slim .`
4. Run your entire compose stack again: `docker compose up`
5. Test it: `curl http://localhost:8080/api/users` should still work.

### 💡 Hints (if stuck):

- The first stage should copy your Go files and run `go build`.
- The second stage should copy the binary **from the first stage** using `COPY --from=builder /app/server .`
- Don't forget to expose a port or set a CMD in the final stage.
- Your final stage doesn't need the Go compiler — just the binary and a minimal OS (Alpine is perfect).
- If you get an error like "command not found: ./server", it might be because Alpine is so minimal it's missing some system libraries. We'll fix that in the next level if needed.

### 🎯 Why this matters:

This is a **pro move**. Real-world Docker deployments almost always use multi-stage builds. You're learning something that actual DevOps engineers do every day. 💪