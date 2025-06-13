// file: pkg/captcha/client.go
package captcha

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// apiURL is the Anti-Captcha endpoint used by the client.
// It can be overridden in tests via SetAPIURL.
var apiURL = "https://api.anti-captcha.com"

// SetAPIURL overrides the Anti-Captcha API URL.
// It is primarily used by tests to point the client at a mock server.
func SetAPIURL(u string) { apiURL = u }

// Client communicates with the Anti-Captcha service.
type Client struct {
	// APIKey is the user's Anti-Captcha API key.
	APIKey string
	// HTTPClient is used to send requests. If nil, http.DefaultClient is used.
	HTTPClient *http.Client
}

// New returns a Client initialised with the provided API key.
func New(apiKey string) *Client {
	return &Client{APIKey: apiKey, HTTPClient: &http.Client{Timeout: 30 * time.Second}}
}

type createResp struct {
	ErrorID          int    `json:"errorId"`
	TaskID           int    `json:"taskId"`
	ErrorCode        string `json:"errorCode"`
	ErrorDescription string `json:"errorDescription"`
}

type resultResp struct {
	ErrorID          int    `json:"errorId"`
	Status           string `json:"status"`
	ErrorCode        string `json:"errorCode"`
	ErrorDescription string `json:"errorDescription"`
	Solution         struct {
		Text               string `json:"text"`
		GRecaptchaResponse string `json:"gRecaptchaResponse"`
	} `json:"solution"`
}

// SolveImage sends an image captcha to the Anti-Captcha service and
// returns the solved text.
func (c *Client) SolveImage(ctx context.Context, b64 string) (string, error) {
	r := map[string]any{
		"clientKey": c.APIKey,
		"task": map[string]string{
			"type": "ImageToTextTask",
			"body": b64,
		},
	}
	var cr createResp
	if err := c.post(ctx, "/createTask", r, &cr); err != nil {
		return "", err
	}
	if cr.ErrorID != 0 {
		return "", fmt.Errorf("createTask %s: %s", cr.ErrorCode, cr.ErrorDescription)
	}
	return c.poll(ctx, cr.TaskID)
}

// SolveRecaptchaV2 solves a reCAPTCHA v2 challenge for the given site.
// The returned string is the token that should be submitted with the form.
func (c *Client) SolveRecaptchaV2(ctx context.Context, siteURL, siteKey string) (string, error) {
	r := map[string]any{
		"clientKey": c.APIKey,
		"task": map[string]string{
			"type":       "NoCaptchaTaskProxyless",
			"websiteURL": siteURL,
			"websiteKey": siteKey,
		},
	}
	var cr createResp
	if err := c.post(ctx, "/createTask", r, &cr); err != nil {
		return "", err
	}
	if cr.ErrorID != 0 {
		return "", fmt.Errorf("createTask %s: %s", cr.ErrorCode, cr.ErrorDescription)
	}
	return c.poll(ctx, cr.TaskID)
}

func (c *Client) post(ctx context.Context, path string, body any, out any) error {
	buf := new(bytes.Buffer)
	if err := json.NewEncoder(buf).Encode(body); err != nil {
		return err
	}
	cli := c.HTTPClient
	if cli == nil {
		cli = http.DefaultClient
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, apiURL+path, buf)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := cli.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("status %d", resp.StatusCode)
	}
	return json.NewDecoder(resp.Body).Decode(out)
}

func (c *Client) poll(ctx context.Context, id int) (string, error) {
	for {
		select {
		case <-ctx.Done():
			return "", ctx.Err()
		case <-time.After(3 * time.Second):
		}
		var rr resultResp
		if err := c.post(ctx, "/getTaskResult", map[string]any{
			"clientKey": c.APIKey,
			"taskId":    id,
		}, &rr); err != nil {
			return "", err
		}
		if rr.ErrorID != 0 {
			return "", fmt.Errorf("getTaskResult %s: %s", rr.ErrorCode, rr.ErrorDescription)
		}
		if rr.Status == "ready" {
			if rr.Solution.Text != "" {
				return rr.Solution.Text, nil
			}
			return rr.Solution.GRecaptchaResponse, nil
		}
	}
}
