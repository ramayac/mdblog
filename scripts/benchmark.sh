#!/usr/bin/env bash
set -euo pipefail

# Check if ApacheBench is installed
if ! command -v ab &> /dev/null; then
  echo "Error: ApacheBench (ab) is not installed. Please install it to run benchmarks."
  echo "On Ubuntu/Debian: sudo apt-get install apache2-utils"
  exit 1
fi

PORT=${PORT:-8089}
export PORT

# Compile the production binary first to ensure we benchmark the latest code
echo "Building bin/mdblog..."
mkdir -p bin
CGO_ENABLED=0 go build -ldflags "-s -w" -o bin/mdblog ./cmd/mdblog

echo "Starting server on port $PORT..."
./bin/mdblog serve > /dev/null 2>&1 &
SERVER_PID=$!

# Ensure the server is terminated on exit
cleanup() {
  echo "Stopping server (PID: $SERVER_PID)..."
  kill "$SERVER_PID" || true
}
trap cleanup EXIT

# Wait for server to start
echo "Waiting for server to start..."
for i in {1..15}; do
  if curl -s "http://localhost:$PORT/" &> /dev/null; then
    break
  fi
  sleep 0.2
done

# Check if server is running
if ! kill -0 "$SERVER_PID" &> /dev/null; then
  echo "Error: Server failed to start."
  exit 1
fi

echo "================================================================="
echo "Benchmarking Home Page (/) - 1000 requests, 10 concurrency"
echo "================================================================="
ab -n 1000 -c 10 "http://localhost:$PORT/"

echo ""
echo "================================================================="
echo "Benchmarking Post Page - 1000 requests, 10 concurrency"
echo "================================================================="
ab -n 1000 -c 10 "http://localhost:$PORT/content/writings/srbyte/srbyte-la-aquitectura-von-neumann"
