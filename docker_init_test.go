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
		"WHISPER_HEALTH_TIMEOUT=0", // Skip health check in tests
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

// TestDockerInitWhisperRetry verifies that docker-init.sh retries on failure
func TestDockerInitWhisperRetry(t *testing.T) {
	tmp := t.TempDir()
	dockerLog := filepath.Join(tmp, "docker.log")
	envFile := filepath.Join(tmp, "env.txt")

	// stub docker binary that fails twice then succeeds
	dockerBin := filepath.Join(tmp, "docker")
	dockerScript := `#!/bin/sh
echo "$@" >>$DOCKER_LOG
if [ "$1" = "ps" ]; then
    exit 0
elif [ "$1" = "run" ]; then
    # Count number of previous run attempts in the log
    attempts=$(grep -c "^run " $DOCKER_LOG || echo 0)
    # Attempt 1 and 2 fail, attempt 3 succeeds
    if [ $attempts -le 2 ]; then
        exit 1  # fail first two attempts
    else
        exit 0  # succeed on third attempt
    fi
else
    exit 0
fi`
	if err := os.WriteFile(dockerBin, []byte(dockerScript), 0755); err != nil {
		t.Fatalf("write docker stub: %v", err)
	}

	// install subtitle-manager stub
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
		"WHISPER_CONTAINER_NAME=test-whisper-retry",
		"WHISPER_PORT=12346",
		"WHISPER_DEVICE=cpu",
		"WHISPER_HEALTH_TIMEOUT=0", // Skip health check in tests
	)

	cmd := exec.Command("sh", "docker-init.sh")
	cmd.Env = env

	if out, err := cmd.CombinedOutput(); err != nil {
		t.Fatalf("docker-init: %v\n%s", err, string(out))
	}

	logData, err := os.ReadFile(dockerLog)
	if err != nil {
		t.Fatalf("read docker log: %v", err)
	}
	
	// Should have attempted 3 times
	runCount := strings.Count(string(logData), "run -d --name test-whisper-retry")
	if runCount != 3 {
		t.Fatalf("expected 3 docker run attempts, got %d: %s", runCount, string(logData))
	}
}

// TestDockerInitWhisperMaxRetries verifies that docker-init.sh exits after max retries
func TestDockerInitWhisperMaxRetries(t *testing.T) {
	tmp := t.TempDir()
	dockerLog := filepath.Join(tmp, "docker.log")

	// stub docker binary that always fails
	dockerBin := filepath.Join(tmp, "docker")
	dockerScript := `#!/bin/sh
echo "$@" >>$DOCKER_LOG
if [ "$1" = "ps" ]; then
    exit 0
elif [ "$1" = "run" ]; then
    exit 1  # always fail
else
    exit 0
fi`
	if err := os.WriteFile(dockerBin, []byte(dockerScript), 0755); err != nil {
		t.Fatalf("write docker stub: %v", err)
	}

	env := append(os.Environ(),
		"PATH="+tmp+":"+os.Getenv("PATH"),
		"DOCKER_LOG="+dockerLog,
		"ENABLE_WHISPER=1",
		"WHISPER_CONTAINER_NAME=test-whisper-fail",
		"WHISPER_DEVICE=cpu",
		"WHISPER_HEALTH_TIMEOUT=0", // Skip health check in tests
	)

	cmd := exec.Command("sh", "docker-init.sh")
	cmd.Env = env

	if err := cmd.Run(); err == nil {
		t.Fatal("expected docker-init to fail after max retries")
	}

	logData, err := os.ReadFile(dockerLog)
	if err != nil {
		t.Fatalf("read docker log: %v", err)
	}
	
	// Should have attempted exactly 3 times
	runCount := strings.Count(string(logData), "run -d --name test-whisper-fail")
	if runCount != 3 {
		t.Fatalf("expected exactly 3 docker run attempts, got %d: %s", runCount, string(logData))
	}
}

// TestDockerInitWhisperContainerExists verifies handling of existing containers
func TestDockerInitWhisperContainerExists(t *testing.T) {
	tmp := t.TempDir()
	dockerLog := filepath.Join(tmp, "docker.log")
	envFile := filepath.Join(tmp, "env.txt")

	// stub docker binary that shows existing running container
	dockerBin := filepath.Join(tmp, "docker")
	dockerScript := `#!/bin/sh
echo "$@" >>$DOCKER_LOG
if [ "$1" = "ps" ] && [ "$3" = "{{.Names}}" ]; then
    if echo "$@" | grep -q -- "--format"; then
        echo "test-whisper-existing"  # container is running
    fi
    exit 0
else
    exit 0
fi`
	if err := os.WriteFile(dockerBin, []byte(dockerScript), 0755); err != nil {
		t.Fatalf("write docker stub: %v", err)
	}

	// install subtitle-manager stub
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
		"WHISPER_CONTAINER_NAME=test-whisper-existing",
		"WHISPER_PORT=12347",
		"WHISPER_DEVICE=cpu",
		"WHISPER_HEALTH_TIMEOUT=0", // Skip health check in tests
	)

	cmd := exec.Command("sh", "docker-init.sh")
	cmd.Env = env

	if out, err := cmd.CombinedOutput(); err != nil {
		t.Fatalf("docker-init: %v\n%s", err, string(out))
	}

	logData, err := os.ReadFile(dockerLog)
	if err != nil {
		t.Fatalf("read docker log: %v", err)
	}
	
	// Should not have attempted to run a new container
	runCount := strings.Count(string(logData), "run -d")
	if runCount > 0 {
		t.Fatalf("should not start new container when one exists, but got %d run attempts: %s", runCount, string(logData))
	}

	// Should still set environment variables
	data, err := os.ReadFile(envFile)
	if err != nil {
		t.Fatalf("read env file: %v", err)
	}
	envContent := string(data)
	if !strings.Contains(envContent, "SM_PROVIDERS_WHISPER_API_URL=http://localhost:12347") {
		t.Fatalf("missing whisper url: %s", envContent)
	}
}

// TestDockerInitWhisperDisabled verifies behavior when ENABLE_WHISPER is not set
func TestDockerInitWhisperDisabled(t *testing.T) {
	tmp := t.TempDir()
	dockerLog := filepath.Join(tmp, "docker.log")
	envFile := filepath.Join(tmp, "env.txt")

	// stub docker binary
	dockerBin := filepath.Join(tmp, "docker")
	if err := os.WriteFile(dockerBin, []byte("#!/bin/sh\necho \"$@\" >>$DOCKER_LOG\nexit 0\n"), 0755); err != nil {
		t.Fatalf("write docker stub: %v", err)
	}

	// install subtitle-manager stub
	smBin := "/usr/local/bin/subtitle-manager"
	if err := os.WriteFile(smBin, []byte("#!/bin/sh\nenv >$ENV_FILE\n"), 0755); err != nil {
		t.Fatalf("install subtitle-manager stub: %v", err)
	}
	t.Cleanup(func() { os.Remove(smBin) })

	env := append(os.Environ(),
		"PATH="+tmp+":"+os.Getenv("PATH"),
		"DOCKER_LOG="+dockerLog,
		"ENV_FILE="+envFile,
		// ENABLE_WHISPER is not set
		"WHISPER_CONTAINER_NAME=test-whisper-disabled",
	)

	cmd := exec.Command("sh", "docker-init.sh")
	cmd.Env = env

	if out, err := cmd.CombinedOutput(); err != nil {
		t.Fatalf("docker-init: %v\n%s", err, string(out))
	}

	// Docker log should be empty - no docker commands executed
	if _, err := os.Stat(dockerLog); err == nil {
		logData, _ := os.ReadFile(dockerLog)
		if len(logData) > 0 {
			t.Fatalf("docker commands executed when ENABLE_WHISPER not set: %s", string(logData))
		}
	}

	// Environment should not have Whisper URLs
	data, err := os.ReadFile(envFile)
	if err != nil {
		t.Fatalf("read env file: %v", err)
	}
	envContent := string(data)
	if strings.Contains(envContent, "SM_PROVIDERS_WHISPER_API_URL") {
		t.Fatalf("whisper url set when disabled: %s", envContent)
	}
}
