package main

import (
	"fmt"
	"github.com/nukosuke/go-zendesk"
	"github.com/nukosuke/go-zendesk/common"
	"net/http"
	"os"
)

func main() {
	httpClient := &http.Client{}
	cred := common.NewAPITokenCredential(os.Getenv("ZD_EMAIL"), os.Getenv("ZD_TOKEN"))

	client, err := zendesk.NewClient(httpClient)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if err = client.SetSubdomain(os.Getenv("ZD_SUBDOMAIN")); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	client.SetCredential(cred)
	triggers, err := client.Core.GetTriggers()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(triggers.Triggers)
}
