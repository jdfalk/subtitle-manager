package cmd

import (
	"context"
	"testing"
)

type fakeWhisperContainer struct {
	started bool
	stopped bool
}

func (f *fakeWhisperContainer) StartContainer(ctx context.Context) error {
	f.started = true
	return nil
}

func (f *fakeWhisperContainer) StopContainer(ctx context.Context) error {
	f.stopped = true
	return nil
}

func (f *fakeWhisperContainer) GetContainerStatus(ctx context.Context) (string, error) {
	return "running", nil
}

func (f *fakeWhisperContainer) IsContainerRunning(ctx context.Context) (bool, error) {
	return true, nil
}

func (f *fakeWhisperContainer) Close() error { return nil }

func TestWhisperStartStopCmd(t *testing.T) {
	fake := &fakeWhisperContainer{}
	orig := newWhisperContainer
	newWhisperContainer = func() (whisperContainer, error) { return fake, nil }
	defer func() { newWhisperContainer = orig }()

	if err := whisperStartCmd.RunE(whisperStartCmd, nil); err != nil {
		t.Fatalf("start: %v", err)
	}
	if !fake.started {
		t.Fatal("StartContainer not called")
	}

	if err := whisperStopCmd.RunE(whisperStopCmd, nil); err != nil {
		t.Fatalf("stop: %v", err)
	}
	if !fake.stopped {
		t.Fatal("StopContainer not called")
	}
}
