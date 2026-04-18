# Stage 1: build imgparse (Rust)
FROM rust:1-bookworm AS rust-builder
WORKDIR /build/imgparse
COPY imgparse/Cargo.toml imgparse/Cargo.lock ./
COPY imgparse/src ./src
COPY imgparse/models ./models
RUN cargo build --release

# Stage 2: build the Go bot
FROM golang:1.22-bookworm AS go-builder
WORKDIR /build/app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o /build/bot .

# Stage 3: runtime
FROM debian:bookworm-slim
RUN apt-get update && apt-get install -y ca-certificates && rm -rf /var/lib/apt/lists/*

COPY --from=rust-builder /build/imgparse/target/release/imgparse /app/imgparse
COPY --from=rust-builder /build/imgparse/models /app/models
COPY --from=go-builder /build/bot /app/bot

VOLUME /data
WORKDIR /data

ENV IMGPARSE_BIN=/app/imgparse
ENV IMGPARSE_MODELS_DIR=/app/models
ENV RESULTS_FILE=/data/wordle_results.json
ENV CURSOR_FILE=/data/cursor.txt

CMD ["/app/bot"]
