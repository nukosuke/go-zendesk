package core

import (
	"net/http"
)

const (
	ZENDESK_DOMAIN = "zendesk.com"
	API_ROOT       = ZENDESK_DOMAIN + "/api/v2"
	API_ENDPOINT   = map[string]string{
		"ticket_form":  API_ROOT + "/ticket_forms",
		"ticket_field": API_ROOT + "/ticket_fields",
	}
)

type CoreAPI struct {
	client *http.Client
	subdomain string
}

func NewClient(client *http.Client, subdomain string) *CoreAPI {
	return &CoreAPI{
		client:    client,
		subdomain: subdomain,
	}
}
