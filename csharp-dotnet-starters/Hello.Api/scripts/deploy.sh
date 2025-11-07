#!/usr/bin/env bash
set -euo pipefail

echo "Deploying to Kubernetes..."

kubectl apply -f manifests/k8s/

