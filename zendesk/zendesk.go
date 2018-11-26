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
	headerRateLimit     = "X-RateLimit-Limit"
	headerRateRemaining = "X-RateLimit-Remaining"
)

var defaultHeaders = map[string]string{
	"User-Agent":   "nukosuke/go-zendesk",
	"Content-Type": "application/json",
}

var subdomainRegexp = regexp.MustCompile("^[a-z][a-z0-9-]+[a-z0-9]$")

// Client of Zendesk API
type Client struct {
	BaseURL    *url.URL
	HTTPClient *http.Client
	Credential Credential
	headers    map[string]string
}

// NewClient creates new Zendesk API client
func NewClient(httpClient *http.Client) (*Client, error) {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	client := &Client{HTTPClient: httpClient}
	client.headers = defaultHeaders
	return client, nil
}

// SetHeader saves HTTP header in client. It will be included all API request
func (z *Client) SetHeader(key string, value string) {
	z.headers[key] = value
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
	z.includeHeaders(req)
	req.SetBasicAuth(z.Credential.Email(), z.Credential.Secret())
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
	z.includeHeaders(req)
	req.SetBasicAuth(z.Credential.Email(), z.Credential.Secret())
	return req, nil
}

// includeHeaders set HTTP headers from client.headers to *http.Request
func (z *Client) includeHeaders(req *http.Request) {
	for key, value := range z.headers {
		req.Header.Set(key, value)
	}
}
