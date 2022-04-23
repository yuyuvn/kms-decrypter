ARG GOLANG_BASE_IMAGE=golang:1.18-alpine
ARG RELEASE_IMAGE=alpine:3.15
FROM $GOLANG_BASE_IMAGE AS builder

WORKDIR /go/src/decrypter
COPY go.* ./
RUN go mod download
COPY . .
RUN go build -o .out/decrypter ./cmd/decrypter/main.go

FROM $RELEASE_IMAGE
COPY --from=builder /go/src/decrypter/.out/decrypter /decrypter
