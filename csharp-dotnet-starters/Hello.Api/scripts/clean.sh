#!/usr/bin/env bash
set -euo pipefail

echo "Cleaning up containers and .NET builds..."
docker compose -f manifests/docker-compose.yaml down -v --remove-orphans
dotnet clean src/Hello.Api

# Delete Kubernetes resources if they exist
K8S_DIR="manifests/k8s"

if kubectl get -f "$K8S_DIR" &> /dev/null; then
    echo "Deleting Kubernetes resources from $K8S_DIR..."
    kubectl delete -f "$K8S_DIR"
else
    echo "No Kubernetes resources found."
fi

echo "✅ Cleanup complete."
