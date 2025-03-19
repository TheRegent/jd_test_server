FROM golang:1.24 as builder
WORKDIR /app
COPY go.mod ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bstrs ./cmd/main.go

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/bstrs .
EXPOSE 8080
CMD ["./bstrs"]
