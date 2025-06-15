package webui

// Smart generate script that only runs npm if dist doesn't exist
//go:generate sh -c "if [ ! -d dist ]; then npm install --legacy-peer-deps && npm run build; else echo 'dist directory exists, skipping npm build'; fi"
