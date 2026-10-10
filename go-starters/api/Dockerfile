FROM golang:1.24 AS builder

ARG OS="linux"
ARG ARCH="amd64"

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY cmd/ cmd/
COPY internal/ internal/
COPY pkg/ pkg/

ENV CGO_ENABLED=1 GOOS=${OS} GOARCH=${ARCH}

RUN go build -o main ./cmd/api/main.go

FROM debian:latest AS final

RUN apt-get update && apt-get install -y libsqlite3-dev sqlite3

WORKDIR /root/

# Copy the compiled binary from the builder stage
COPY --from=builder /app/main .

RUN chmod +x main

# Command to run the executable
CMD ["./main"]
