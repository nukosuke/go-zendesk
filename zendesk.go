package zendesk

import (
	"fmt"
	"github.com/zenform/go-zendesk/common"
	"github.com/zenform/go-zendesk/core"
	"net/http"
	"net/url"
)

const (
	baseURLFormat       = "https://%s.zendesk.com/api/v2"
	userAgent           = "zenform/go-zendesk"
	headerRateLimit     = "X-RateLimit-Limit"
	headerRateRemaining = "X-RateLimit-Remaining"
)

// Client of Zendesk API
type Client struct {
	Core *core.Service
	//TODO: support other APIs
}

// NewClient creates new Zendesk API client
func NewClient(httpClient *http.Client, subdomain string) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	baseURLString := fmt.Sprintf(baseURLFormat, subdomain)
	baseURL, err := url.Parse(baseURLString)
	if err != nil {
		return nil
	}

	client := &Client{}
	service := &common.Service{
		HTTPClient: httpClient,
		BaseURL:    baseURL,
		UserAgent:  userAgent,
	}

	client.Core = (*core.Service)(service)
	// other services...
	return client
}
