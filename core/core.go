package core

import (
	"net/http"
)

type CoreAPI struct {
	client *http.Client
}

func NewClient(client *http.Client) *CoreAPI {
	return &CoreAPI{
		client: client,
	}
}
