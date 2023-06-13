package main

import (
	"fmt"

	"github.com/stripe/stripe-go/v74"
	"github.com/stripe/stripe-go/v74/customer"
	"github.com/stripe/stripe-go/v74/paymentmethod"
)

func main() {
	stripe.Key = "sk_test_sample"
	sampleCustomerID := "cus_sample"
	stripeParams := stripe.PaymentMethodParams{
		Card: &stripe.PaymentMethodCardParams{
			Number:   stripe.String("4242424242424242"),
			ExpMonth: stripe.Int64(12),
			ExpYear:  stripe.Int64(2024),
			CVC:      stripe.String("424"),
		},
		Type: stripe.String(string(stripe.PaymentMethodTypeCard)),
	}

	pm, err := paymentmethod.New(&stripeParams)
	if err != nil {
		panic(err)
	}

	_, err = paymentmethod.Attach(
		pm.ID,
		&stripe.PaymentMethodAttachParams{
			Customer: stripe.String(sampleCustomerID),
		},
	)
	if err != nil {
		panic(err)
	}

	_, err = customer.Update("cus_O3WsRj7vAEkKyA", &stripe.CustomerParams{
		InvoiceSettings: &stripe.CustomerInvoiceSettingsParams{
			DefaultPaymentMethod: &pm.ID,
		},
	})
	if err != nil {
		panic(err)
	}

	fmt.Println(pm.ID)

}
