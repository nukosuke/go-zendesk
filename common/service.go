package common

import (
	"net/http"
	"net/url"
)

type Service struct {
	HTTPClient    *http.Client
	BaseURL   *url.URL
	UserAgent string
}
