package csvtaxrule

import (
	"fmt"

	"github.com/autopilot3/chargebee-go"
	"github.com/autopilot3/chargebee-go/models/csvtaxrule"
)

func Create(params *csvtaxrule.CreateRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/csv_tax_rules"), params)
}
