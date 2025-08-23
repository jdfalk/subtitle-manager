// file: pkg/authserver/server.go
// version: 1.1.0
// guid: d14bd9a1-4b55-4b57-bffc-4adc165ebd1a

package authserver

import (
	"context"
	"database/sql"
	"strconv"
	"time"

	authpb "github.com/jdfalk/gcommon/sdks/go/v1/common"
	gauth "github.com/jdfalk/subtitle-manager/pkg/gcommonauth"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// Server implements authpb.AuthServiceServer using the local database
// and gcommonauth helpers.
type Server struct {
	DB *sql.DB
	authpb.UnimplementedAuthServiceServer
}

// NewServer creates a new Server instance.
func NewServer(db *sql.DB) *Server {
	return &Server{DB: db}
}

// Authenticate validates user credentials and issues a session token.
func (s *Server) Authenticate(ctx context.Context, req *authpb.AuthenticateRequest) (*authpb.AuthenticateResponse, error) {
	var userID int64
	switch creds := req.Credentials.(type) {
	case *authpb.AuthenticateRequest_Password:
		id, err := gauth.AuthenticateUser(s.DB, creds.Password.Username, creds.Password.Password)
		if err != nil {
			return nil, err
		}
		userID = id
	case *authpb.AuthenticateRequest_ApiKey:
		id, err := gauth.ValidateAPIKey(s.DB, creds.ApiKey.Key)
		if err != nil {
			return nil, err
		}
		userID = id
	default:
		return nil, sql.ErrNoRows
	}

	token, err := gauth.GenerateSession(s.DB, userID, 24*time.Hour)
	if err != nil {
		return nil, err
	}

	return &authpb.AuthenticateResponse{
		AccessToken: token,
		TokenType:   "session",
		ExpiresIn:   int32(24 * 60 * 60),
		UserInfo: &authpb.UserInfo{
			Id:       strconv.FormatInt(userID, 10),
			Username: credsUsername(req),
		},
	}, nil
}

func credsUsername(req *authpb.AuthenticateRequest) string {
	if p, ok := req.Credentials.(*authpb.AuthenticateRequest_Password); ok {
		return p.Password.Username
	}
	return ""
}

// ValidateToken verifies a session token and returns expiration info.
func (s *Server) ValidateToken(ctx context.Context, req *authpb.ValidateTokenRequest) (*authpb.ValidateTokenResponse, error) {
	var userID int64
	var expires time.Time
	row := s.DB.QueryRow(`SELECT user_id, expires_at FROM sessions WHERE token = ?`, req.AccessToken)
	if err := row.Scan(&userID, &expires); err != nil {
		if err == sql.ErrNoRows {
			return &authpb.ValidateTokenResponse{Valid: false}, nil
		}
		return nil, err
	}
	if time.Now().After(expires) {
		return &authpb.ValidateTokenResponse{Valid: false}, nil
	}

	return &authpb.ValidateTokenResponse{
		Valid:     true,
		Subject:   strconv.FormatInt(userID, 10),
		ExpiresAt: timestamppb.New(expires),
		ExpiresIn: int32(time.Until(expires).Seconds()),
	}, nil
}
