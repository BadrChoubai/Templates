FROM mcr.microsoft.com/dotnet/aspnet:9.0 AS base

WORKDIR /app
EXPOSE 8080
EXPOSE 8081

FROM mcr.microsoft.com/dotnet/sdk:9.0 AS build
ARG BUILD_CONFIGURATION=Release
# Copy the entire project (Dockerfile is inside Hello.Api), restore and build
WORKDIR /src
COPY . .
# Restore and build the project file in the current directory
RUN dotnet restore "Hello.Api.csproj"
RUN dotnet build "Hello.Api.csproj" -c $BUILD_CONFIGURATION -o /app/build

FROM build AS publish
ARG BUILD_CONFIGURATION=Release
RUN dotnet publish "Hello.Api.csproj" -c $BUILD_CONFIGURATION -o /app/publish /p:UseAppHost=false

FROM base AS final
WORKDIR /app
COPY --from=publish /app/publish .
## Install curl (not present in the aspnet runtime image) so the HEALTHCHECK can
## probe the HTTP endpoint. This increases image size slightly; if you prefer a
## zero-footprint probe, consider using an external health-probe sidecar or the
## orchestrator's native HTTP checks.
RUN apt-get update \
	&& apt-get install -y --no-install-recommends curl \
	&& rm -rf /var/lib/apt/lists/*

# Healthcheck probes the application's health endpoint on port 8080 where the
# app listens inside the container. Adjust if your app uses a different port.

HEALTHCHECK --interval=30s --timeout=5s --start-period=5s --retries=3 \
    CMD curl -f http://localhost:8080/health || exit 1

ENTRYPOINT ["dotnet", "Hello.Api.dll"]
