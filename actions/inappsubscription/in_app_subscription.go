package inappsubscription

import (
	"fmt"
	"net/url"

	"github.com/autopilot3/chargebee-go"
	"github.com/autopilot3/chargebee-go/models/inappsubscription"
)

func ProcessReceipt(id string, params *inappsubscription.ProcessReceiptRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/in_app_subscriptions/%v/process_purchase_command", url.PathEscape(id)), params)
}
func ImportReceipt(id string, params *inappsubscription.ImportReceiptRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/in_app_subscriptions/%v/import_receipt", url.PathEscape(id)), params)
}
func ImportSubscription(id string, params *inappsubscription.ImportSubscriptionRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/in_app_subscriptions/%v/import_subscription", url.PathEscape(id)), params)
}
func RetrieveStoreSubs(id string, params *inappsubscription.RetrieveStoreSubsRequestParams) chargebee.RequestObj {
	return chargebee.Send("POST", fmt.Sprintf("/in_app_subscriptions/%v/retrieve", url.PathEscape(id)), params)
}
