package chargebee

import (
	"github.com/autopilot3/chargebee-go/models/addon"
	"github.com/autopilot3/chargebee-go/models/address"
	"github.com/autopilot3/chargebee-go/models/advanceinvoiceschedule"
	"github.com/autopilot3/chargebee-go/models/attacheditem"
	"github.com/autopilot3/chargebee-go/models/card"
	"github.com/autopilot3/chargebee-go/models/comment"
	"github.com/autopilot3/chargebee-go/models/contact"
	"github.com/autopilot3/chargebee-go/models/contractterm"
	"github.com/autopilot3/chargebee-go/models/coupon"
	"github.com/autopilot3/chargebee-go/models/couponcode"
	"github.com/autopilot3/chargebee-go/models/couponset"
	"github.com/autopilot3/chargebee-go/models/creditnote"
	"github.com/autopilot3/chargebee-go/models/customer"
	"github.com/autopilot3/chargebee-go/models/differentialprice"
	"github.com/autopilot3/chargebee-go/models/download"
	"github.com/autopilot3/chargebee-go/models/entitlementoverride"
	"github.com/autopilot3/chargebee-go/models/estimate"
	"github.com/autopilot3/chargebee-go/models/event"
	"github.com/autopilot3/chargebee-go/models/export"
	"github.com/autopilot3/chargebee-go/models/feature"
	"github.com/autopilot3/chargebee-go/models/gift"
	"github.com/autopilot3/chargebee-go/models/hierarchy"
	"github.com/autopilot3/chargebee-go/models/hostedpage"
	"github.com/autopilot3/chargebee-go/models/invoice"
	"github.com/autopilot3/chargebee-go/models/item"
	"github.com/autopilot3/chargebee-go/models/itementitlement"
	"github.com/autopilot3/chargebee-go/models/itemfamily"
	"github.com/autopilot3/chargebee-go/models/itemprice"
	"github.com/autopilot3/chargebee-go/models/order"
	"github.com/autopilot3/chargebee-go/models/paymentintent"
	"github.com/autopilot3/chargebee-go/models/paymentsource"
	"github.com/autopilot3/chargebee-go/models/plan"
	"github.com/autopilot3/chargebee-go/models/portalsession"
	"github.com/autopilot3/chargebee-go/models/promotionalcredit"
	"github.com/autopilot3/chargebee-go/models/quote"
	"github.com/autopilot3/chargebee-go/models/quotedcharge"
	"github.com/autopilot3/chargebee-go/models/quotedsubscription"
	"github.com/autopilot3/chargebee-go/models/quotelinegroup"
	"github.com/autopilot3/chargebee-go/models/resourcemigration"
	"github.com/autopilot3/chargebee-go/models/sitemigrationdetail"
	"github.com/autopilot3/chargebee-go/models/subscription"
	"github.com/autopilot3/chargebee-go/models/subscriptionentitlement"
	"github.com/autopilot3/chargebee-go/models/taxwithheld"
	"github.com/autopilot3/chargebee-go/models/thirdpartypaymentmethod"
	"github.com/autopilot3/chargebee-go/models/timemachine"
	"github.com/autopilot3/chargebee-go/models/token"
	"github.com/autopilot3/chargebee-go/models/transaction"
	"github.com/autopilot3/chargebee-go/models/unbilledcharge"
	"github.com/autopilot3/chargebee-go/models/usage"
	"github.com/autopilot3/chargebee-go/models/virtualbankaccount"
)

type ResultList struct {
	List       []*Result `json:"list"`
	NextOffset string    `json:"next_offset"`
}
type Result struct {
	Subscription            *subscription.Subscription                       `json:"subscription,omitempty"`
	ContractTerm            *contractterm.ContractTerm                       `json:"contract_term,omitempty"`
	AdvanceInvoiceSchedule  *advanceinvoiceschedule.AdvanceInvoiceSchedule   `json:"advance_invoice_schedule,omitempty"`
	Customer                *customer.Customer                               `json:"customer,omitempty"`
	Hierarchy               *hierarchy.Hierarchy                             `json:"hierarchy,omitempty"`
	Contact                 *contact.Contact                                 `json:"contact,omitempty"`
	Token                   *token.Token                                     `json:"token,omitempty"`
	PaymentSource           *paymentsource.PaymentSource                     `json:"payment_source,omitempty"`
	ThirdPartyPaymentMethod *thirdpartypaymentmethod.ThirdPartyPaymentMethod `json:"third_party_payment_method,omitempty"`
	VirtualBankAccount      *virtualbankaccount.VirtualBankAccount           `json:"virtual_bank_account,omitempty"`
	Card                    *card.Card                                       `json:"card,omitempty"`
	PromotionalCredit       *promotionalcredit.PromotionalCredit             `json:"promotional_credit,omitempty"`
	Invoice                 *invoice.Invoice                                 `json:"invoice,omitempty"`
	TaxWithheld             *taxwithheld.TaxWithheld                         `json:"tax_withheld,omitempty"`
	CreditNote              *creditnote.CreditNote                           `json:"credit_note,omitempty"`
	UnbilledCharge          *unbilledcharge.UnbilledCharge                   `json:"unbilled_charge,omitempty"`
	Order                   *order.Order                                     `json:"order,omitempty"`
	Gift                    *gift.Gift                                       `json:"gift,omitempty"`
	Transaction             *transaction.Transaction                         `json:"transaction,omitempty"`
	HostedPage              *hostedpage.HostedPage                           `json:"hosted_page,omitempty"`
	Estimate                *estimate.Estimate                               `json:"estimate,omitempty"`
	Quote                   *quote.Quote                                     `json:"quote,omitempty"`
	QuotedSubscription      *quotedsubscription.QuotedSubscription           `json:"quoted_subscription,omitempty"`
	QuotedCharge            *quotedcharge.QuotedCharge                       `json:"quoted_charge,omitempty"`
	QuoteLineGroup          *quotelinegroup.QuoteLineGroup                   `json:"quote_line_group,omitempty"`
	Plan                    *plan.Plan                                       `json:"plan,omitempty"`
	Addon                   *addon.Addon                                     `json:"addon,omitempty"`
	Coupon                  *coupon.Coupon                                   `json:"coupon,omitempty"`
	CouponSet               *couponset.CouponSet                             `json:"coupon_set,omitempty"`
	CouponCode              *couponcode.CouponCode                           `json:"coupon_code,omitempty"`
	Address                 *address.Address                                 `json:"address,omitempty"`
	Usage                   *usage.Usage                                     `json:"usage,omitempty"`
	Event                   *event.Event                                     `json:"event,omitempty"`
	Comment                 *comment.Comment                                 `json:"comment,omitempty"`
	Download                *download.Download                               `json:"download,omitempty"`
	PortalSession           *portalsession.PortalSession                     `json:"portal_session,omitempty"`
	SiteMigrationDetail     *sitemigrationdetail.SiteMigrationDetail         `json:"site_migration_detail,omitempty"`
	ResourceMigration       *resourcemigration.ResourceMigration             `json:"resource_migration,omitempty"`
	TimeMachine             *timemachine.TimeMachine                         `json:"time_machine,omitempty"`
	Export                  *export.Export                                   `json:"export,omitempty"`
	PaymentIntent           *paymentintent.PaymentIntent                     `json:"payment_intent,omitempty"`
	ItemFamily              *itemfamily.ItemFamily                           `json:"item_family,omitempty"`
	Item                    *item.Item                                       `json:"item,omitempty"`
	ItemPrice               *itemprice.ItemPrice                             `json:"item_price,omitempty"`
	AttachedItem            *attacheditem.AttachedItem                       `json:"attached_item,omitempty"`
	DifferentialPrice       *differentialprice.DifferentialPrice             `json:"differential_price,omitempty"`
	Feature                 *feature.Feature                                 `json:"feature,omitempty"`
	SubscriptionEntitlement *subscriptionentitlement.SubscriptionEntitlement `json:"subscription_entitlement,omitempty"`
	ItemEntitlement         *itementitlement.ItemEntitlement                 `json:"item_entitlement,omitempty"`
	EntitlementOverride     *entitlementoverride.EntitlementOverride         `json:"entitlement_override,omitempty"`
	UnbilledCharges         []*unbilledcharge.UnbilledCharge                 `json:"unbilled_charges,omitempty"`
	CreditNotes             []*creditnote.CreditNote                         `json:"credit_notes,omitempty"`
	AdvanceInvoiceSchedules []*advanceinvoiceschedule.AdvanceInvoiceSchedule `json:"advance_invoice_schedules,omitempty"`
	Hierarchies             []*hierarchy.Hierarchy                           `json:"hierarchies,omitempty"`
	Downloads               []*download.Download                             `json:"downloads,omitempty"`
	Invoices                []*invoice.Invoice                               `json:"invoices,omitempty"`
	DifferentialPrices      []*differentialprice.DifferentialPrice           `json:"differential_prices,omitempty"`
}
