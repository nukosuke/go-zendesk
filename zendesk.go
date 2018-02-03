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

func (c *Client) SetCredential(cred *common.Credential) error {
	subdomain := cred.Subdomain
	if !subdomainRegexp.MatchString(subdomain) {
		return fmt.Errorf("%s is invalid subdomain", subdomain)
	}

	baseURLString := fmt.Sprintf(baseURLFormat, subdomain)
	baseURL, err := url.Parse(baseURLString)
	if err != nil {
		return err
	}

	c.Core.BaseURL = baseURL
	c.Core.Credential = cred
	return nil
}
