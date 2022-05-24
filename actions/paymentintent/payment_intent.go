package paymentintent

import (
	"fmt"
	"github.com/autopilot3/chargebee-go"
	"github.com/autopilot3/chargebee-go/models/paymentintent"
	"net/url"
)

func Create(params *paymentintent.CreateRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/payment_intents"), params)
}
func Update(id string, params *paymentintent.UpdateRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/payment_intents/%v", url.PathEscape(id)), params)
}
func Retrieve(id string) chargebee.RequestObj {
	return chargebee.Send("GET", fmt.Sprintf("/payment_intents/%v", url.PathEscape(id)), nil)
}
