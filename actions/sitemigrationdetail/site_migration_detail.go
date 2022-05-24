package sitemigrationdetail

import (
	"fmt"
	"github.com/autopilot3/chargebee-go"
	"github.com/autopilot3/chargebee-go/models/sitemigrationdetail"
)

func List(params *sitemigrationdetail.ListRequestParams) chargebee.RequestObj {
	return chargebee.SendList("GET", fmt.Sprintf("/site_migration_details"), params)
}
