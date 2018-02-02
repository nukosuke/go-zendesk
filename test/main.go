package main

import (
	"fmt"
	"github.com/zenform/go-zendesk"
	"net/http"
	"os"
)

func main() {
	httpClient := &http.Client{}
	client, err := zendesk.NewClient(httpClient, os.Getenv("ZD_SUBDOMAIN"))
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
