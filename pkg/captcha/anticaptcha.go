// file: pkg/captcha/anticaptcha.go
package captcha

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// Solver resolves captchas required by some providers.
type Solver interface {
	// Solve returns the captcha solution token for the given site key and page URL.
	Solve(ctx context.Context, siteKey, pageURL string) (string, error)
}

// TwoCaptcha implements the Solver interface using the 2captcha service.
type TwoCaptcha struct {
	APIKey  string
	BaseURL string
	client  *http.Client
}

// NewTwoCaptcha creates a solver with the provided API key.
func NewTwoCaptcha(apiKey string) *TwoCaptcha {
	return &TwoCaptcha{APIKey: apiKey, BaseURL: "https://2captcha.com", client: &http.Client{Timeout: 60 * time.Second}}
}

// Solve submits the captcha and polls until it is solved.
func (t *TwoCaptcha) Solve(ctx context.Context, siteKey, pageURL string) (string, error) {
	form := url.Values{
		"key":       {t.APIKey},
		"method":    {"userrecaptcha"},
		"googlekey": {siteKey},
		"pageurl":   {pageURL},
		"json":      {"1"},
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, t.BaseURL+"/in.php", strings.NewReader(form.Encode()))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	resp, err := t.client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	var r struct {
		Status  int    `json:"status"`
		Request string `json:"request"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&r); err != nil {
		return "", err
	}
	if r.Status != 1 {
		return "", errors.New(r.Request)
	}
	id := r.Request
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()
	for {
		select {
		case <-ctx.Done():
			return "", ctx.Err()
		case <-ticker.C:
			u := t.BaseURL + "/res.php?key=" + url.QueryEscape(t.APIKey) + "&action=get&id=" + id + "&json=1"
			req2, err := http.NewRequestWithContext(ctx, http.MethodGet, u, nil)
			if err != nil {
				return "", err
			}
			resp2, err := t.client.Do(req2)
			if err != nil {
				return "", err
			}
			defer resp2.Body.Close()
			var r2 struct {
				Status  int    `json:"status"`
				Request string `json:"request"`
			}
			if err := json.NewDecoder(resp2.Body).Decode(&r2); err != nil {
				return "", err
			}
			if r2.Status == 1 {
				return r2.Request, nil
			}
			if r2.Request != "CAPCHA_NOT_READY" {
				return "", errors.New(r2.Request)
			}
		}
	}
}
