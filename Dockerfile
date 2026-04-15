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

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Copy built frontend from bun stage
COPY --from=client-builder /app/client/dist ./client/dist

RUN go build -o app


# Step 3: Runtime
FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/app .
COPY --from=builder /app/client/dist ./client/dist

CMD ["./app"]
