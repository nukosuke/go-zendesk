package core

import (
	"net/http"
)

const (
	ZendeskDomain = "zendesk.com"
	APIRoot       = ZendeskDomain + "/api/v2"
)

var APIEndpoint = map[string]string{
	"ticket_form":  APIRoot + "/ticket_forms",
	"ticket_field": APIRoot + "/ticket_fields",
	"triggers":     APIRoot + "/triggers",
}

type CoreAPI struct {
	client    *http.Client
	subdomain string
}

// create new CoreAPI
func NewClient(client *http.Client, subdomain string) *CoreAPI {
	return &CoreAPI{
		client:    client,
		subdomain: subdomain,
	}
}
