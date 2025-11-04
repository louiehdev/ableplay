FROM golang:alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o ableplay_web ./cmd/frontend

FROM debian:bookworm-slim
WORKDIR /app
COPY --from=builder /app/ableplay_web .
EXPOSE 9633
CMD ["./ableplay_web"]