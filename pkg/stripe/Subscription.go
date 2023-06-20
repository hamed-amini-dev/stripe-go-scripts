package stripe

import (
	"github.com/stripe/stripe-go/v74"
	"github.com/stripe/stripe-go/v74/invoice"
	"github.com/stripe/stripe-go/v74/subscription"
)

// https://github.com/stripe-samples/placing-a-hold/blob/main/without-webhooks/server/php/index.php
func CreateSubscription(promotionCode, customerID, price string, trialPeriodDays, qty int64, payout bool) (*stripe.Subscription, error) {
	var PromotionCode *string
	if promotionCode == "" {
		PromotionCode = nil
	} else {
		PromotionCode = stripe.String(promotionCode)
	}

	sParams := &stripe.SubscriptionParams{
		Customer:        &customerID,
		TrialPeriodDays: stripe.Int64(trialPeriodDays),
		Items: []*stripe.SubscriptionItemsParams{
			{
				Price:    &price,
				Quantity: stripe.Int64(qty),
			},
		},
		CollectionMethod: stripe.String(string(stripe.InvoiceCollectionMethodSendInvoice)),
		DaysUntilDue:     stripe.Int64(0),
		PromotionCode:    PromotionCode,
	}

	stripeSubscription, err := subscription.New(sParams)
	if err != nil {
		return nil, err
	}

	if payout {
		//pay invoice without paying
		_, err = invoice.Pay(stripeSubscription.LatestInvoice.ID, &stripe.InvoicePayParams{
			PaidOutOfBand: stripe.Bool(true),
		})

		if err != nil {
			return nil, err
		}

		stripeSubscription, err = subscription.Get(stripeSubscription.ID, nil)
		if err != nil {
			return nil, err
		}
	}

	return stripeSubscription, nil

}
