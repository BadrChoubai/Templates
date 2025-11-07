#!/usr/bin/env bash
set -euo pipefail

echo "Starting API, Prometheus, and Grafana..."
docker compose -f manifests/docker-compose.yaml up --build
