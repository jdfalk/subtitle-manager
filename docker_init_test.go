package main

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

// TestDockerInitWhisper verifies docker-init.sh launches a whisper container
// and exports expected environment variables when ENABLE_WHISPER=1.
func TestDockerInitWhisper(t *testing.T) {
	tmp := t.TempDir()
	dockerLog := filepath.Join(tmp, "docker.log")
	envFile := filepath.Join(tmp, "env.txt")

	// stub docker binary
	dockerBin := filepath.Join(tmp, "docker")
	if err := os.WriteFile(dockerBin, []byte("#!/bin/sh\necho \"$@\" >>$DOCKER_LOG\n[ \"$1\" = ps ] && exit 0\nexit 0\n"), 0755); err != nil {
		t.Fatalf("write docker stub: %v", err)
	}

	// install subtitle-manager stub to /usr/local/bin
	smBin := "/usr/local/bin/subtitle-manager"
	if err := os.WriteFile(smBin, []byte("#!/bin/sh\nenv >$ENV_FILE\n"), 0755); err != nil {
		t.Fatalf("install subtitle-manager stub: %v", err)
	}
	t.Cleanup(func() { os.Remove(smBin) })

	env := append(os.Environ(),
		"PATH="+tmp+":"+os.Getenv("PATH"),
		"DOCKER_LOG="+dockerLog,
		"ENV_FILE="+envFile,
		"ENABLE_WHISPER=1",
		"WHISPER_CONTAINER_NAME=test-whisper",
		"WHISPER_PORT=12345",
		"WHISPER_DEVICE=cpu",
	)

	cmd := exec.Command("sh", "docker-init.sh")
	cmd.Env = env

	if out, err := cmd.CombinedOutput(); err != nil {
		t.Fatalf("docker-init: %v\n%s", err, string(out))
	}

	data, err := os.ReadFile(envFile)
	if err != nil {
		t.Fatalf("read env file: %v", err)
	}
	envContent := string(data)
	if !strings.Contains(envContent, "SM_PROVIDERS_WHISPER_API_URL=http://localhost:12345") {
		t.Fatalf("missing whisper url: %s", envContent)
	}
	if !strings.Contains(envContent, "SM_OPENAI_API_URL=http://localhost:12345/v1") {
		t.Fatalf("missing openai url: %s", envContent)
	}

	logData, err := os.ReadFile(dockerLog)
	if err != nil {
		t.Fatalf("read docker log: %v", err)
	}
	if !strings.Contains(string(logData), "run -d --name test-whisper") {
		t.Fatalf("docker run not executed: %s", string(logData))
	}
}
