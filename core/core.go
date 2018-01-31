package core

import (
	"net/http"
)

const (
	// ZendeskDomain is domain of zendesk
	ZendeskDomain = "zendesk.com"

	// APIRoot is API path including version prefix
	APIRoot       = ZendeskDomain + "/api/v2"
)

// APIEndpoint is resource name and API endpoint mapping
var APIEndpoint = map[string]string{
	"ticket_form":  APIRoot + "/ticket_forms",
	"ticket_field": APIRoot + "/ticket_fields",
	"triggers":     APIRoot + "/triggers",
}

type CoreAPI struct {
	client    *http.Client
	subdomain string
}

// NewClient create new CoreAPI
func NewClient(client *http.Client, subdomain string) *CoreAPI {
	return &CoreAPI{
		client:    client,
		subdomain: subdomain,
	}
}
