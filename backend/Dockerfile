FROM golang:1.20-alpine as builder
WORKDIR /app
COPY . .
RUN go mod download && go build -o server ./cmd

FROM alpine:3.16
WORKDIR /app
COPY --from=builder /app/server .
EXPOSE 8080
CMD ["./server"]