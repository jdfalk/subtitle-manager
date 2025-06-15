package webui

// Smart generate script that only runs npm if dist doesn't exist and we're not in Docker
//go:generate sh -c "if [ ! -d dist ] && [ -z \"$DOCKER_BUILD\" ]; then npm install --legacy-peer-deps && npm run build; else echo 'dist directory exists or Docker build detected, skipping npm build'; fi"
