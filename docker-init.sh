#!/bin/sh
# file: docker-init.sh
# Launch optional Whisper ASR container and then start Subtitle Manager
set -e

WHISPER_CONTAINER_NAME=${WHISPER_CONTAINER_NAME:-whisper-asr-service}
WHISPER_IMAGE=${WHISPER_IMAGE:-onerahmet/openai-whisper-asr-webservice:latest}
WHISPER_PORT=${WHISPER_PORT:-9000}

cleanup() {
  if [ "$ENABLE_WHISPER" = "1" ]; then
    docker stop "$WHISPER_CONTAINER_NAME" >/dev/null 2>&1 || true
  fi
}
trap cleanup EXIT

if [ "$ENABLE_WHISPER" = "1" ]; then
  if command -v docker >/dev/null 2>&1; then
    if ! docker ps --format '{{.Names}}' | grep -q "^$WHISPER_CONTAINER_NAME$"; then
      echo "Starting Whisper ASR container..."
      docker run -d --name "$WHISPER_CONTAINER_NAME" \
        --gpus all \
        -p ${WHISPER_PORT}:9000 \
        -e ASR_MODEL=${WHISPER_MODEL:-base} \
        -e ASR_DEVICE=${WHISPER_DEVICE:-cuda} \
        "$WHISPER_IMAGE" >/dev/null
    fi
    export SM_PROVIDERS_WHISPER_API_URL=${SM_PROVIDERS_WHISPER_API_URL:-http://localhost:${WHISPER_PORT}}
  else
    echo "Docker not available; cannot launch Whisper ASR service" >&2
  fi
fi

exec /usr/local/bin/subtitle-manager "$@"
