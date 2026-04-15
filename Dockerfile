FROM golang:1.25 AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o app

# Step 2: Run
FROM alpine:latest

WORKDIR /app
COPY --from=builder /app/app .
COPY client/dist /app/client/dist

CMD ["./app"]
