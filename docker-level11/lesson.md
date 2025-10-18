# 🎯 Level 11 – The Image Slim-Down 📦

## 🎓 Lesson 11 – Multi-Stage Builds

Here's a problem you might not have noticed yet:

Your current Go Dockerfile looks like this:

```dockerfile
FROM golang:1.25.1
WORKDIR /app
COPY . .
RUN go build -o server .
CMD ["./server"]
```

This works, but it's **fat**. The `golang:1.25.1` image is **~800MB**. You're shipping the entire Go compiler, all build tools, source code — everything — into production. 

But here's the thing: **you only need the compiled binary to run**. You don't need Go installed in production.

That's where **multi-stage builds** come in. It's like:

1. **Stage 1 (Builder)**: Use the big Go image, compile your app → produces a tiny binary.
2. **Stage 2 (Runtime)**: Use a minimal base image (like `alpine` or `scratch`), copy only the binary from Stage 1 → done.

Result? Your image goes from **~800MB → ~20MB** (or less). 🚀

### Example:

```dockerfile
# Stage 1: Build
FROM golang:1.25.1 AS builder
WORKDIR /app
COPY . .
RUN go build -o server .

# Stage 2: Runtime
FROM alpine:latest
COPY --from=builder /app/server .
CMD ["./server"]
```

Notice:
- `AS builder` → names the first stage.
- `--from=builder` → copies files **from** that stage, not from your host.
- Final image only has the binary + Alpine OS. Tiny.

### Why care?

- **Faster deployments** (smaller images = faster push/pull to registries).
- **Less storage** on servers.
- **Smaller attack surface** (fewer packages = fewer vulnerabilities).
- **Better for Kubernetes** (orchestrating thousands of containers is painful with bloated images).