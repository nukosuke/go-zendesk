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
		Subdomain: os.Getenv("ZD_SUBDOMAIN"),
		AuthType:  common.APIToken,
		Email:     os.Getenv("ZD_EMAIL"),
		APIToken:  os.Getenv("ZD_TOKEN"),
	}

	client, err := zendesk.NewClient(httpClient, cred)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	//client.SetAPIToken(os.Getenv("ZD_EMAIL"), os.Getenv("ZD_TOKEN"))

	triggers, err := client.Core.GetTriggers()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(triggers.Triggers)
}
