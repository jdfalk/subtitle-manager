// file: pkg/providers/generic/generic.go
package generic

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"path/filepath"
	"time"

	"github.com/spf13/viper"
)

// Client implements the providers.Provider interface using a configurable
// HTTP endpoint.
type Client struct {
	APIURL   string
	Username string
	Password string
	APIKey   string

	HTTPClient *http.Client
}

// New returns a Client configured from Viper keys under
// "providers.generic".
func New() *Client {
	return &Client{
		APIURL:     viper.GetString("providers.generic.api_url"),
		Username:   viper.GetString("providers.generic.username"),
		Password:   viper.GetString("providers.generic.password"),
		APIKey:     viper.GetString("providers.generic.api_key"),
		HTTPClient: &http.Client{Timeout: 15 * time.Second},
	}
}

// Fetch retrieves a subtitle using the configured endpoint. Query parameters
// include the file name, language and optional credentials.
func (c *Client) Fetch(ctx context.Context, mediaPath, lang string) ([]byte, error) {
	if c.APIURL == "" {
		return nil, fmt.Errorf("api url not configured")
	}
	u, err := url.Parse(c.APIURL)
	if err != nil {
		return nil, err
	}
	q := u.Query()
	q.Set("file", filepath.Base(mediaPath))
	q.Set("lang", lang)
	if c.Username != "" {
		q.Set("username", c.Username)
	}
	if c.Password != "" {
		q.Set("password", c.Password)
	}
	if c.APIKey != "" {
		q.Set("api_key", c.APIKey)
	}
	u.RawQuery = q.Encode()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, err
	}
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status %d", resp.StatusCode)
	}
	return io.ReadAll(resp.Body)
}
