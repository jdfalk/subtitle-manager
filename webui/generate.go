package webui

// Smart generate script that only runs npm if built assets are missing and we're not in Docker
//go:generate sh -c "if [ ! -d dist/assets ] && [ -z \"$DOCKER_BUILD\" ]; then npm install --legacy-peer-deps && npm run build; else echo 'dist assets exist or Docker build detected, skipping npm build'; fi"
