package auth

import (
	"testing"
	"time"

	"subtitle-manager/pkg/database"
)

func TestSetUserRole(t *testing.T) {
	db, err := database.Open(":memory:")
	if err != nil {
		t.Fatalf("open: %v", err)
	}
	defer db.Close()
	if err := CreateUser(db, "u", "p", "", "user"); err != nil {
		t.Fatalf("create: %v", err)
	}
	if err := SetUserRole(db, "u", "admin"); err != nil {
		t.Fatalf("set role: %v", err)
	}
	ok, err := CheckPermission(db, 1, "all")
	if err != nil {
		t.Fatalf("check: %v", err)
	}
	if !ok {
		t.Fatal("permission not granted")
	}
}

func TestOneTimeToken(t *testing.T) {
	db, err := database.Open(":memory:")
	if err != nil {
		t.Fatalf("open: %v", err)
	}
	defer db.Close()
	if err := CreateUser(db, "u", "p", "e@example.com", "user"); err != nil {
		t.Fatalf("create: %v", err)
	}
	token, err := GenerateOneTimeToken(db, 1, time.Minute)
	if err != nil {
		t.Fatalf("gen token: %v", err)
	}
	id, err := ConsumeOneTimeToken(db, token)
	if err != nil {
		t.Fatalf("consume: %v", err)
	}
	if id != 1 {
		t.Fatalf("unexpected user id %d", id)
	}
	if _, err := ConsumeOneTimeToken(db, token); err == nil {
		t.Fatal("expected second consume to fail")
	}
}
