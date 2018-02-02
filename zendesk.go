package zendesk

import (
	"fmt"
	"github.com/zenform/go-zendesk/common"
	"github.com/zenform/go-zendesk/core"
	"net/http"
	"net/url"
	"regexp"
)

const (
	baseURLFormat       = "https://%s.zendesk.com/api/v2"
	userAgent           = "zenform/go-zendesk"
	headerRateLimit     = "X-RateLimit-Limit"
	headerRateRemaining = "X-RateLimit-Remaining"
)

var subdomainRegexp = regexp.MustCompile("^[a-z][a-z0-9]+$")

// Client of Zendesk API
type Client struct {
	Core *core.Service
	//TODO: support other APIs
}

// NewClient creates new Zendesk API client
func NewClient(httpClient *http.Client, subdomain string) (*Client, error) {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	if !subdomainRegexp.MatchString(subdomain) {
		return nil, fmt.Errorf("%s is invalid subdomain", subdomain)
	}

	baseURLString := fmt.Sprintf(baseURLFormat, subdomain)
	baseURL, err := url.Parse(baseURLString)
	if err != nil {
		return nil, err
	}

	client := &Client{}
	service := &common.Service{
		HTTPClient: httpClient,
		BaseURL:    baseURL,
		UserAgent:  userAgent,
	}

	client.Core = (*core.Service)(service)
	// other services...
	return client, nil
}
