package main

import (
	"fmt"
	"github.com/nukosuke/go-zendesk/zendesk"
	"net/http"
	"os"
)

func main() {
	httpClient := &http.Client{}
	cred := zendesk.NewAPITokenCredential(os.Getenv("ZD_EMAIL"), os.Getenv("ZD_TOKEN"))

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
	triggers, err := client.GetTriggers()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(triggers.Triggers)

	ticketFields, page, err := client.GetTicketFields()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(ticketFields)

	ticketForms, err := client.GetTicketForms()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("===== page has next")
	fmt.Println(page.HasNext())
	fmt.Println(ticketForms.TicketForms)

	client.PostTicketField(zendesk.TicketField{Type: "text", Title: "Age"})
}
