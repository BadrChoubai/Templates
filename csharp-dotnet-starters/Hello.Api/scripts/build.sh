#!/usr/bin/env bash
set -euo pipefail

PROJECT="src/Hello.Api"
IMAGE_NAME=${IMAGE_NAME:-hello-api}
REGISTRY=${REGISTRY:-localhost:5000}

echo "🔨 Building .NET project..."
dotnet build "$PROJECT" -c Release

echo "🐳 Building Docker image..."
docker build -t "${REGISTRY}/${IMAGE_NAME}:latest" -f "${PROJECT}/Dockerfile" "${PROJECT}"
