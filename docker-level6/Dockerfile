FROM golang:1.25.1
WORKDIR /app
COPY . .
RUN go build -o server .
CMD ["./server"]
