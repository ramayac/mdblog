#!/usr/bin/env bash
set -euo pipefail

PORT=${PORT:-8091}
export PORT

if ! command -v node &> /dev/null; then
  echo "Error: node is required to run layout lint."
  exit 1
fi

echo "Building bin/mdblog and content artifacts..."
mkdir -p bin
CGO_ENABLED=0 go build -ldflags "-s -w" -o bin/mdblog ./cmd/mdblog
./bin/mdblog build > /dev/null

echo "Starting server on port $PORT..."
./bin/mdblog serve > /dev/null 2>&1 &
SERVER_PID=$!

cleanup() {
  kill "$SERVER_PID" 2>/dev/null || true
}
trap cleanup EXIT

for _ in {1..20}; do
  if curl -sf "http://localhost:$PORT/" > /dev/null 2>&1; then
    break
  fi
  sleep 0.2
done

if ! kill -0 "$SERVER_PID" 2>/dev/null; then
  echo "Error: Server failed to start."
  exit 1
fi

node scripts/lint-layout.mjs "http://localhost:$PORT"
