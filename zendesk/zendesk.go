package zendesk

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"regexp"
	"strings"
)

const (
	baseURLFormat       = "https://%s.zendesk.com/api/v2"
	userAgent           = "nukosuke/go-zendesk"
	headerRateLimit     = "X-RateLimit-Limit"
	headerRateRemaining = "X-RateLimit-Remaining"
)

var subdomainRegexp = regexp.MustCompile("^[a-z][a-z0-9]+$")

// Client of Zendesk API
type Client struct {
	BaseURL    *url.URL
	UserAgent  string
	HTTPClient *http.Client
	Credential Credential
}

// NewClient creates new Zendesk API client
func NewClient(httpClient *http.Client) (*Client, error) {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	client := &Client{HTTPClient: httpClient}
	return client, nil
}

// SetSubdomain saves subdomain in client. It will be used
// when call API
func (z *Client) SetSubdomain(subdomain string) error {
	if !subdomainRegexp.MatchString(subdomain) {
		return fmt.Errorf("%s is invalid subdomain", subdomain)
	}

	baseURLString := fmt.Sprintf(baseURLFormat, subdomain)
	baseURL, err := url.Parse(baseURLString)
	if err != nil {
		return err
	}

	z.BaseURL = baseURL
	return nil
}

// SetCredential saves credential in client. It will be set
// to request header when call API
func (z *Client) SetCredential(cred Credential) {
	z.Credential = cred
}

// NewGetRequest create GET *http.Request with headers which are required for authentication.
func (z Client) NewGetRequest(path string) (*http.Request, error) {
	req, err := http.NewRequest("GET", z.BaseURL.String()+path, nil)
	if err != nil {
		return nil, err
	}
	req.SetBasicAuth(z.Credential.Email(), z.Credential.Secret())
	req.Header.Set("User-Agent", z.UserAgent)
	return req, nil
}

// NewPostRequest create POST *http.Request with headers which are required for authentication.
func (z Client) NewPostRequest(path string, payload interface{}) (*http.Request, error) {
	bytes, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", z.BaseURL.String()+path, strings.NewReader(string(bytes)))
	if err != nil {
		return nil, err
	}
	req.SetBasicAuth(z.Credential.Email(), z.Credential.Secret())
	req.Header.Set("User-Agent", z.UserAgent)
	req.Header.Set("Content-Type", "application/json")
	return req, nil
}
