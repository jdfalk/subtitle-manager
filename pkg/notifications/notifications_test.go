package notifications

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/smtp"
	"testing"
)

func TestDiscordNotifier(t *testing.T) {
	var got string
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		var m map[string]string
		if err := json.NewDecoder(r.Body).Decode(&m); err != nil {
			t.Fatalf("failed to decode request body: %v", err)
		}
		got = m["content"]
	}))
	defer srv.Close()

	n := DiscordNotifier{WebhookURL: srv.URL, Client: srv.Client()}
	if err := n.Notify(context.Background(), "hello"); err != nil {
		t.Fatalf("notify: %v", err)
	}
	if got != "hello" {
		t.Fatalf("unexpected message: %s", got)
	}
}

func TestTelegramNotifier(t *testing.T) {
	var got string
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		var m map[string]string
		if err := json.NewDecoder(r.Body).Decode(&m); err != nil {
			t.Fatalf("failed to decode request body: %v", err)
		}
		got = m["text"]
	}))
	defer srv.Close()

	n := TelegramNotifier{BotToken: "t", ChatID: "1", Client: srv.Client(), APIBase: srv.URL}
	if err := n.Notify(context.Background(), "hi"); err != nil {
		t.Fatalf("notify: %v", err)
	}
	if got != "hi" {
		t.Fatalf("unexpected message: %s", got)
	}
}

func TestSMTPNotifier(t *testing.T) {
	var addr, from string
	var to []string
	var body []byte
	n := SMTPNotifier{
		Addr: "smtp:25",
		From: "a@example.com",
		To:   []string{"b@example.com"},
		Send: func(a string, _ smtp.Auth, f string, t []string, msg []byte) error {
			addr = a
			from = f
			to = t
			body = msg
			return nil
		},
	}
	if err := n.Notify(context.Background(), "hi there"); err != nil {
		t.Fatalf("notify: %v", err)
	}
	if addr != "smtp:25" || from != "a@example.com" || len(to) != 1 || to[0] != "b@example.com" {
		t.Fatalf("unexpected params")
	}
	if string(body) != "Subject: Subtitle Manager\r\n\r\nhi there" {
		t.Fatalf("unexpected body: %s", body)
	}
}
