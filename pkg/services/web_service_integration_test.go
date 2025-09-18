// file: pkg/services/web_service_integration_test.go
// version: 1.0.0
// guid: e2d3c4b5-a6f7-4890-b1c2-d3e4f5a6b7c8

package services

import (
    "context"
    "net"
    "testing"

    webv1 "github.com/jdfalk/subtitle-manager/pkg/web/v1"
    "github.com/stretchr/testify/require"
    "google.golang.org/grpc"
    "google.golang.org/grpc/test/bufconn"
)

const bufSize = 1024 * 1024

func startBufconnServer(t *testing.T) (*grpc.ClientConn, func()) {
    t.Helper()
    lis := bufconn.Listen(bufSize)
    s := grpc.NewServer()
    webv1.RegisterWebServiceServer(s, NewWebService())

    go func() {
        if err := s.Serve(lis); err != nil {
            // test helper: do not fail test here, client side will surface errors
        }
    }()

    dialer := func(context.Context, string) (net.Conn, error) { return lis.Dial() }
    ctx := context.Background()
    conn, err := grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(dialer), grpc.WithInsecure())
    require.NoError(t, err)

    cleanup := func() {
        conn.Close()
        s.Stop()
        lis.Close()
    }
    return conn, cleanup
}

func TestWebService_Bufconn_LogoutAndHealth(t *testing.T) {
    conn, cleanup := startBufconnServer(t)
    defer cleanup()

    client := webv1.NewWebServiceClient(conn)
    ctx := context.Background()

    // HealthCheck should succeed
    _, err := client.HealthCheck(ctx, &webv1.HealthCheckRequest{})
    require.NoError(t, err)

    // LogoutUser should return success=true
    out, err := client.LogoutUser(ctx, &webv1.LogoutUserRequest{})
    require.NoError(t, err)
    require.NotNil(t, out)
    require.True(t, out.GetSuccess())
}
