package main

import (
	alerts "github.com/opsgenie/opsgenie-go-sdk/alerts"
	ogcli "github.com/opsgenie/opsgenie-go-sdk/client"
	"fmt"
)

var API_KEY string = "YOUR API KEY HERE"

func main() {

	cli := new (ogcli.OpsGenieClient)
	cli.SetApiKey(API_KEY)

	alertCli, cliErr := cli.Alert()

	if cliErr != nil {
		panic(cliErr)
	}

	// list the alerts
	listreq := alerts.ListAlertsRequest{}
	listresp, listErr := alertCli.List(listreq)

	if listErr != nil {
		panic(listErr)
	}

	for _, alert := range listresp.Alerts {
		fmt.Println("Id: ", alert.Id)
		fmt.Println("Alias: ", alert.Alias)
		fmt.Println("Message: ", alert.Message)
		fmt.Println("Status: ", alert.Status)
		fmt.Println("IsSeen?: ", alert.IsSeen)
		fmt.Println("Acknowledged?: ", alert.Acknowledged)
		fmt.Println("Created at: ", alert.CreatedAt)
		fmt.Println("Updated at: ", alert.UpdatedAt)
		fmt.Println("Tiny id: ", alert.TinyId)
	}
}