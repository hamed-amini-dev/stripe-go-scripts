package main

import (
	"fmt"

	"github.com/stripe/stripe-go/v74"
	"github.com/stripe/stripe-go/v74/invoice"
	"github.com/stripe/stripe-go/v74/subscription"
)

// https://github.com/stripe-samples/placing-a-hold/blob/main/without-webhooks/server/php/index.php
func Create(promotionCode, customerID, price string, trialPeriodDays, qty int64, payout bool) *stripe.Subscription {
	var PromotionCode *string
	if promotionCode == "" {
		PromotionCode = nil
	} else {
		PromotionCode = stripe.String(promotionCode)
	}

	/* plan, err := s.plan.Get(params.PlanID)
	if err != nil {
		return nil, err
	} */

	/* if params.Qty <= 0 {
		params.Qty = 1
	} */

	sParams := &stripe.SubscriptionParams{
		Customer:        &customerID,
		TrialPeriodDays: stripe.Int64(trialPeriodDays),
		Items: []*stripe.SubscriptionItemsParams{
			{
				// Price:    stripe.String(plan.GetRecurringPriceID()),
				Price:    &price,
				Quantity: stripe.Int64(qty),
			},
		},
		CollectionMethod: stripe.String(string(stripe.InvoiceCollectionMethodSendInvoice)),
		DaysUntilDue:     stripe.Int64(0),
		PromotionCode:    PromotionCode,
	}
	// sParams.AddMetadata("plan_id", plan.ID)
	/* if params.DefaultPaymentMethodID != "" {
		cParams.DefaultPaymentMethod = stripe.String(params.DefaultPaymentMethodID)
	} */
	stripeSubscription, err := subscription.New(sParams)
	if err != nil {
		panic(err)
	}

	if payout {
		//pay invoice without paying
		_, err = invoice.Pay(stripeSubscription.LatestInvoice.ID, &stripe.InvoicePayParams{
			PaidOutOfBand: stripe.Bool(true),
		})

		if err != nil {
			panic(err)
		}

		stripeSubscription, err = subscription.Get(stripeSubscription.ID, nil)
		if err != nil {
			panic(err)
		}
	}

}

func main() {
	stripe.Key = "sk_test_51LfLkJLdKtGMf7ZsI7yvCPZizcWoRxYSqN49SCsqzuQv3UVwu7t4hrvOS4CSakkpZPnKkV3sjICOENjK0MEZWAUT00u7sKul49"

	fmt.Println(customerID.ID)
	//cus_O3WsRj7vAEkKyA

}
