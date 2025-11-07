#!/usr/bin/env bash
set -euo pipefail

IMAGE_NAME=${IMAGE_NAME:-hello-api}
REGISTRY=${REGISTRY:-localhost:5000}

echo "📦 Pushing image ${REGISTRY}/${IMAGE_NAME}:latest..."
docker push "${REGISTRY}/${IMAGE_NAME}:latest"

