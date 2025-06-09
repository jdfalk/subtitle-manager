package auth

import (
	"subtitle-manager/pkg/database"
	"testing"
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
