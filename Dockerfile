# Step 1: Build client (Bun + Vite)
FROM oven/bun:latest AS client-builder

WORKDIR /app/client

COPY client/package.json client/bun.lockb* ./
RUN bun install

COPY client/ .

RUN bunx vite build


# Step 2: Build Go app
FROM golang:1.25 AS builder

WORKDIR /app

# C toolchain + sqlite for CGO
RUN apt-get update && apt-get install -y \
    gcc \
    sqlite3 \
    libsqlite3-dev \
    && rm -rf /var/lib/apt/lists/*

COPY go.mod go.sum ./
RUN go mod download

COPY . .

COPY --from=client-builder /app/client/dist ./client/dist

# ARM64 Linux build with CGO enabled
RUN CGO_ENABLED=1 GOOS=linux GOARCH=arm64 go build -o app


# Step 3: Runtime (IMPORTANT: NOT Alpine)
FROM debian:bookworm-slim

WORKDIR /app

# SQLite runtime library
RUN apt-get update && apt-get install -y \
    sqlite3 \
    libsqlite3-0 \
    && rm -rf /var/lib/apt/lists/*

COPY --from=builder /app/app .
COPY --from=builder /app/client/dist ./client/dist

CMD ["./app"]
