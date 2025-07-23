// file: pkg/authserver/server_test.go
// version: 1.0.0
// guid: c5c8d260-2641-45dc-80d2-4dbb941bdb7e

package authserver

import (
	"context"
	"testing"

	authpb "github.com/jdfalk/gcommon/pkg/auth/proto"
	gauth "github.com/jdfalk/subtitle-manager/pkg/gcommonauth"
	"github.com/jdfalk/subtitle-manager/pkg/testutil"
)

// TestAuthenticatePassword verifies password-based authentication
func TestAuthenticatePassword(t *testing.T) {
	db := testutil.GetTestDB(t)
	defer db.Close()

	testutil.MustNoError(t, "create user", gauth.CreateUser(db, "u1", "pass", "e@example.com", "user"))

	srv := NewServer(db)
	resp, err := srv.Authenticate(context.Background(), &authpb.AuthenticateRequest{
		Credentials: &authpb.AuthenticateRequest_Password{Password: &authpb.PasswordCredentials{
			Username: "u1",
			Password: "pass",
		}},
	})
	testutil.MustNoError(t, "authenticate", err)
	testutil.MustNotEqual(t, "token", "", resp.AccessToken)

	val, err := srv.ValidateToken(context.Background(), &authpb.ValidateTokenRequest{AccessToken: resp.AccessToken})
	testutil.MustNoError(t, "validate", err)
	testutil.MustEqual(t, "valid", true, val.Valid)
}

// TestAuthenticateAPIKey verifies API key authentication
func TestAuthenticateAPIKey(t *testing.T) {
	db := testutil.GetTestDB(t)
	defer db.Close()

	testutil.MustNoError(t, "create user", gauth.CreateUser(db, "u2", "pass", "", "user"))
	key, err := gauth.GenerateAPIKey(db, 1)
	testutil.MustNoError(t, "generate key", err)

	srv := NewServer(db)
	resp, err := srv.Authenticate(context.Background(), &authpb.AuthenticateRequest{
		Credentials: &authpb.AuthenticateRequest_ApiKey{ApiKey: &authpb.APIKeyCredentials{Key: key}},
	})
	testutil.MustNoError(t, "authenticate", err)
	testutil.MustNotEqual(t, "token", "", resp.AccessToken)
}
