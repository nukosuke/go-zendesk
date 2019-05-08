package zendesk

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
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

// Error an error type containing the http response from zendesk
type Error struct {
	body []byte
	resp *http.Response
}

// Error the error string for this error
func (e Error) Error() string {
	msg := string(e.body)
	if msg == "" {
		msg = http.StatusText(e.Status())
	}

	return fmt.Sprintf("%d: %s", e.resp.StatusCode, msg)
}

// Body is the Body of the HTTP response
func (e Error) Body() io.ReadCloser {
	return ioutil.NopCloser(bytes.NewBuffer(e.body))
}

// Headers the HTTP headers returned from zendesk
func (e Error) Headers() http.Header {
	return e.resp.Header
}

// Status the HTTP status code returned from zendesk
func (e Error) Status() int {
	return e.resp.StatusCode
}

var subdomainRegexp = regexp.MustCompile("^[a-z][a-z0-9-]+[a-z0-9]$")

// Client of Zendesk API
type Client struct {
	baseURL    *url.URL
	httpClient *http.Client
	credential Credential
	headers    map[string]string
}

// NewClient creates new Zendesk API client
func NewClient(httpClient *http.Client) (*Client, error) {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	client := &Client{httpClient: httpClient}
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

	z.baseURL = baseURL
	return nil
}

// SetEndpointURL replace full URL of endpoint without subdomain validation.
// This is mainly used for testing to point to mock API server.
func (z *Client) SetEndpointURL(newURL string) error {
	baseURL, err := url.Parse(newURL)
	if err != nil {
		return err
	}

	z.baseURL = baseURL
	return nil
}

// SetCredential saves credential in client. It will be set
// to request header when call API
func (z *Client) SetCredential(cred Credential) {
	z.credential = cred
}

// Get get JSON data from API and returns its body as []bytes
func (z *Client) Get(path string) ([]byte, error) {
	req, err := http.NewRequest(http.MethodGet, z.baseURL.String()+path, nil)
	if err != nil {
		return nil, err
	}

	z.prepareRequest(req)

	resp, err := z.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, Error{
			body: body,
			resp: resp,
		}
	}
	return body, nil
}

// Post send data to API and returns response body as []bytes
func (z *Client) Post(path string, data interface{}) ([]byte, error) {
	bytes, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, z.baseURL.String()+path, strings.NewReader(string(bytes)))
	if err != nil {
		return nil, err
	}
	z.prepareRequest(req)

	resp, err := z.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusCreated {
		return nil, Error{
			body: body,
			resp: resp,
		}
	}

	return body, nil
}

// Put sends data to API and returns response body as []bytes
func (z *Client) Put(path string, data interface{}) ([]byte, error) {
	bytes, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPut, z.baseURL.String()+path, strings.NewReader(string(bytes)))
	if err != nil {
		return nil, err
	}
	z.prepareRequest(req)

	resp, err := z.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, Error{
			body: body,
			resp: resp,
		}
	}

	return body, nil
}

// Delete sends data to API and returns an error if unsuccessful
func (z *Client) Delete(path string) error {
	req, err := http.NewRequest(http.MethodDelete, z.baseURL.String()+path, nil)
	if err != nil {
		return err
	}
	z.prepareRequest(req)

	resp, err := z.httpClient.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusNoContent {
		return Error{
			body: body,
			resp: resp,
		}
	}

	return nil
}

// prepare request sets common request variables such as authn and user agent
func (z *Client) prepareRequest(req *http.Request) {
	z.includeHeaders(req)
	req.SetBasicAuth(z.credential.Email(), z.credential.Secret())
}

// includeHeaders set HTTP headers from client.headers to *http.Request
func (z *Client) includeHeaders(req *http.Request) {
	for key, value := range z.headers {
		req.Header.Set(key, value)
	}
}
