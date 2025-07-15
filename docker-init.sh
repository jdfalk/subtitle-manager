#!/bin/sh
# file: docker-init.sh
# Launch optional Whisper ASR container and then start Subtitle Manager
set -e

WHISPER_CONTAINER_NAME=${WHISPER_CONTAINER_NAME:-whisper-asr-service}
WHISPER_IMAGE=${WHISPER_IMAGE:-onerahmet/openai-whisper-asr-webservice:latest}
WHISPER_PORT=${WHISPER_PORT:-9000}
# Max number of startup attempts before failing
WHISPER_MAX_RETRIES=${WHISPER_MAX_RETRIES:-3}

# Validate Whisper model and device selections
case "${WHISPER_MODEL:-base}" in
  tiny|base|small|medium|large) ;;
  *)
    echo "Invalid WHISPER_MODEL, defaulting to 'base'" >&2
    WHISPER_MODEL=base
    ;;
esac

case "${WHISPER_DEVICE:-cuda}" in
  cpu|cuda) ;;
  *)
    echo "Invalid WHISPER_DEVICE, defaulting to 'cuda'" >&2
    WHISPER_DEVICE=cuda
    ;;
esac

cleanup() {
  if [ "$ENABLE_WHISPER" = "1" ]; then
    docker stop "$WHISPER_CONTAINER_NAME" >/dev/null 2>&1 || true
  fi
}
trap cleanup EXIT

if [ "$ENABLE_WHISPER" = "1" ]; then
  if command -v docker >/dev/null 2>&1; then
    # Check if container already exists
    if docker ps -a --format '{{.Names}}' | grep -q "^$WHISPER_CONTAINER_NAME$"; then
      if docker ps --format '{{.Names}}' | grep -q "^$WHISPER_CONTAINER_NAME$"; then
        echo "Whisper ASR container is already running"
      else
        echo "Removing stopped Whisper ASR container..."
        docker rm "$WHISPER_CONTAINER_NAME" >/dev/null 2>&1 || true
      fi
    fi
    
    # Start container if not running
    if ! docker ps --format '{{.Names}}' | grep -q "^$WHISPER_CONTAINER_NAME$"; then
      echo "Starting Whisper ASR container..."
      gpu_flag=""
      if [ "${WHISPER_DEVICE:-cuda}" != "cpu" ]; then
        gpu_flag="--gpus all"
      fi
      
      # Start container with retry logic
      retry_count=0
      max_retries=${WHISPER_MAX_RETRIES:-3}
      container_started=0
      
      while [ "$retry_count" -lt "$max_retries" ] && [ "$container_started" -eq 0 ]; do
        if docker run -d --name "$WHISPER_CONTAINER_NAME" "$gpu_flag" \
          -p "${WHISPER_PORT}":9000 \
          -e ASR_MODEL="${WHISPER_MODEL:-base}" \
          -e ASR_DEVICE="${WHISPER_DEVICE:-cuda}" \
          "$WHISPER_IMAGE" >/dev/null 2>&1; then
          echo "Whisper ASR container started successfully"
          container_started=1
        else
          retry_count=$((retry_count + 1))
          echo "Failed to start Whisper container (attempt $retry_count/$max_retries)"
          if [ "$retry_count" -lt "$max_retries" ]; then
            sleep 2
            docker rm "$WHISPER_CONTAINER_NAME" >/dev/null 2>&1 || true
          fi
        fi
      done
      
      if [ $container_started -eq 0 ]; then
        echo "Failed to start Whisper ASR container after $max_retries attempts" >&2
        exit 1
      fi
      
      # Basic readiness check with configurable timeout
      max_wait=${WHISPER_HEALTH_TIMEOUT:-10}
      if [ "$max_wait" -gt 0 ]; then
        echo "Waiting for Whisper ASR service to be ready..."
        ready_count=0

        while [ "$ready_count" -lt "$max_wait" ]; do
          # Check if container is still running (basic health check)
          if docker ps --format '{{.Names}}' | grep -q "^$WHISPER_CONTAINER_NAME$"; then
            # Try HTTP health check if tools are available
            health_ok=0
            if command -v curl >/dev/null 2>&1; then
              if curl -s -f "http://localhost:${WHISPER_PORT}/health" >/dev/null 2>&1; then
                health_ok=1
              fi
            elif command -v wget >/dev/null 2>&1; then
              if wget -q --spider "http://localhost:${WHISPER_PORT}/health" >/dev/null 2>&1; then
                health_ok=1
              fi
            fi
            
            if [ $health_ok -eq 1 ]; then
              echo "Whisper ASR service is ready"
              break
            fi
          else
            echo "Whisper ASR container stopped unexpectedly" >&2
            exit 1
          fi
          
          ready_count=$((ready_count + 1))
          if [ "$ready_count" -ge "$max_wait" ]; then
            echo "Warning: Whisper ASR service may not be ready after ${max_wait}s" >&2
            break
          fi
          sleep 1
        done
      fi
    fi
    
    export SM_PROVIDERS_WHISPER_API_URL="${SM_PROVIDERS_WHISPER_API_URL:-http://localhost:${WHISPER_PORT}}"
    export SM_OPENAI_API_URL="${SM_OPENAI_API_URL:-http://localhost:${WHISPER_PORT}/v1}"
  else
    echo "Docker not available; cannot launch Whisper ASR service" >&2
  fi
fi

exec /usr/local/bin/subtitle-manager "$@"
