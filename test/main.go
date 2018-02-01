package main

import (
	"net/http"
	"os"
	"fmt"
	"github.com/zenform/go-zendesk"
)

func main() {
	httpClient := &http.Client{}
	client := zendesk.NewClient(httpClient, os.Getenv("ZD_SUBDOMAIN"))
	//client.SetAPIToken(os.Getenv("ZD_EMAIL"), os.Getenv("ZD_TOKEN"))

	triggers, err := client.Core.GetTriggers()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(triggers.Triggers)
}
