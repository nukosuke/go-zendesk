package main

import (
	"fmt"
	"github.com/zenform/go-zendesk"
	"github.com/zenform/go-zendesk/common"
	"net/http"
	"os"
)

func main() {
	httpClient := &http.Client{}
	cred := &common.Credential{
		AuthType:  common.APIToken,
		Email:     os.Getenv("ZD_EMAIL"),
		APIToken:  os.Getenv("ZD_TOKEN"),
	}

	client, err := zendesk.NewClient(httpClient)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if err = client.SetSubdomain(os.Getenv("ZD_SUBDOMAIN")); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if err = client.SetCredential(cred); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	triggers, err := client.Core.GetTriggers()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(triggers.Triggers)
}
