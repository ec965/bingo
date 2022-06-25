FROM golang:1.18-alpine as builder
WORKDIR /app
COPY go.mod ./
COPY go.sum ./
COPY pkgs ./pkgs
COPY main.go ./
RUN go build main.go

FROM alpine:3.14.4
WORKDIR /app
COPY --from=builder /app/main .
CMD ["./main"]
