# Hello API

A minimal ASP.NET Core API with OpenTelemetry instrumentation, health checks, and Docker Compose setup for local observability.

---

## 🧠 Overview

**Hello API** exposes a simple greeting endpoint (`/`) and health checks under `/health`, `/health/live`, and `/health/ready`.  
It includes built-in telemetry for tracing, metrics, and logging using OpenTelemetry.

### Features

- **ASP.NET Core Minimal API**
- **OpenTelemetry** (metrics + tracing)
- **Prometheus & Grafana integration**
- **Health checks for Kubernetes**
- **Environment-based configuration**
- **Docker Compose for local observability stack**

---

## 🧰 Running the App Locally

### Prerequisites

- [.NET 8 SDK](https://dotnet.microsoft.com/en-us/download)
- [Docker](https://www.docker.com/)
- (Optional) [Docker Compose](https://docs.docker.com/compose/)

---

### Option 1: Run directly with `dotnet`

```bash
cd src/Hello.Api
dotnet run
```

Access the app at: `http://localhost:5014`

Health check:

```bash
curl http://localhost:5014/health
```

---

### Option 2: Run with Docker Compose

To start the API, Prometheus, and Grafana together:

```bash
docker compose -f manifests/docker-compose.yml up --build
```

Access the services:

- **Hello API:** [http://localhost:8080](http://localhost:8080)
- **Prometheus:** [http://localhost:9090](http://localhost:9090)
- **Grafana:** [http://localhost:3000](http://localhost:3000)
  - Use default credentials: `admin` / `admin`

You can stop the stack with:

```bash
docker compose -f manifests/docker-compose.yml down
```

### Running with Kubernetes

You can deploy the application to a local Kubernetes cluster (like Minikube or Kind)
using the provided manifests in the `manifests/k8s` directory.

---

## Scripts

- `./scripts/build.sh` - Builds the Docker image for the API.
- `./scripts/run.sh` - Runs the API locally using `docker compose`.
- `./scripts/push.sh` - Pushes the Docker image to a container registry.
- `./scripts/deploy.sh` - Deploys the API to a Kubernetes cluster.
- `./scripts/clean.sh` - Cleans up .NET, Docker images and Kubernetes deployment.

---
