ARG GO_VERSION=1.26

FROM golang:${GO_VERSION}-bookworm AS builder

ARG OS="linux"
ARG ARCH="amd64"

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY cmd/ cmd/
COPY internal/ internal/
COPY pkg/ pkg/

# go-sqlite3 requires cgo; it bundles SQLite, so no system library is needed at runtime.
ENV CGO_ENABLED=1 GOOS=${OS} GOARCH=${ARCH}

RUN go build -o main ./cmd/api

FROM debian:bookworm-slim AS final

WORKDIR /root/

# SQLite creates the database file but not its directory.
RUN mkdir -p .db

# Copy the compiled binary from the builder stage
COPY --from=builder /app/main .

EXPOSE 8084

# Command to run the executable
CMD ["./main"]
