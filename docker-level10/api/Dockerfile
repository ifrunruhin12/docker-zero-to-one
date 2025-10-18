FROM golang:1.25.1 AS builder
WORKDIR /app
COPY . .
RUN go build -o server .

FROM alpine:latest
COPY --from=builder /app/server .
CMD ["./server"]

