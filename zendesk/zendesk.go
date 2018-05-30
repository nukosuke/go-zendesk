package zendesk

import (
	"fmt"
	"github.com/nukosuke/go-zendesk/common"
	"github.com/nukosuke/go-zendesk/core"
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
func NewClient(httpClient *http.Client) (*Client, error) {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	client := &Client{}
	service := &common.Service{
		HTTPClient: httpClient,
		UserAgent:  userAgent,
	}

	client.Core = (*core.Service)(service)
	// other services...
	return client, nil
}

// SetSubdomain saves subdomain in client. It will be used
// when call API
func (c *Client) SetSubdomain(subdomain string) error {
	if !subdomainRegexp.MatchString(subdomain) {
		return fmt.Errorf("%s is invalid subdomain", subdomain)
	}

	baseURLString := fmt.Sprintf(baseURLFormat, subdomain)
	baseURL, err := url.Parse(baseURLString)
	if err != nil {
		return err
	}

	c.Core.BaseURL = baseURL
	return nil
}

// SetCredential saves credential in client. It will be set
// to request header when call API
func (c *Client) SetCredential(cred common.Credential) {
	c.Core.Credential = cred
}
