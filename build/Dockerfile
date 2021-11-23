FROM golang:1.17-alpine AS builder

WORKDIR /go/src/decrypter
COPY go.* ./
RUN go mod download
COPY . .
RUN go build -o .out/decrypter ./cmd/decrypter/main.go

FROM alpine:3.14
COPY --from=builder /go/src/decrypter/.out/decrypter /decrypter
