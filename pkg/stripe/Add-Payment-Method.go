package stripe

import (
	"github.com/stripe/stripe-go/v74"
	"github.com/stripe/stripe-go/v74/customer"
	"github.com/stripe/stripe-go/v74/paymentmethod"
)

func AddPaymentMethod(customerID string, number, cvc string, expMpnth, expYear int64) (string, error) {
	stripeParams := stripe.PaymentMethodParams{
		Card: &stripe.PaymentMethodCardParams{
			Number:   stripe.String(number),
			ExpMonth: stripe.Int64(expMpnth),
			ExpYear:  stripe.Int64(expYear),
			CVC:      stripe.String(cvc),
		},
		Type: stripe.String(string(stripe.PaymentMethodTypeCard)),
	}

	pm, err := paymentmethod.New(&stripeParams)
	if err != nil {
		return "", err
	}

	_, err = paymentmethod.Attach(
		pm.ID,
		&stripe.PaymentMethodAttachParams{
			Customer: stripe.String(customerID),
		},
	)
	if err != nil {
		return "", err
	}

	_, err = customer.Update(customerID, &stripe.CustomerParams{
		InvoiceSettings: &stripe.CustomerInvoiceSettingsParams{
			DefaultPaymentMethod: &pm.ID,
		},
	})
	if err != nil {
		return "", err
	}

	return pm.ID, nil
}
