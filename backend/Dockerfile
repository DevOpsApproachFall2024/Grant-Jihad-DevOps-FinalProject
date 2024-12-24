FROM golang:1.23-alpine as builder
WORKDIR /app
COPY . .
RUN go mod tidy && go build -o server ./cmd

FROM alpine:3.16
WORKDIR /app
COPY --from=builder /app/server .
EXPOSE 8080
CMD ["./server"]
