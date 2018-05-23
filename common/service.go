package common

import (
	"net/http"
	"net/url"
)

// Service is base struct
type Service struct {
	HTTPClient *http.Client
	BaseURL    *url.URL
	UserAgent  string
	Credential Credential
}
