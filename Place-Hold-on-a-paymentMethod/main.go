package main

import (
	"errors"

	"github.com/stripe/stripe-go/v74"
	"github.com/stripe/stripe-go/v74/customer"
	"github.com/stripe/stripe-go/v74/paymentintent"
	"github.com/stripe/stripe-go/v74/paymentmethod"
)

func PaymentMethodsList(customerID string) ([]*stripe.PaymentMethod, error) {
	params := &stripe.PaymentMethodListParams{
		Customer: stripe.String(customerID),
		Type:     stripe.String("card"),
	}

	cardList := paymentmethod.List(params).PaymentMethodList().Data

	return cardList, nil
}

func GetDefaultPaymentMethod(customerID string) (string, error) {
	cus, err := customer.Get(customerID, nil)
	if err != nil {
		return "", err
	}
	paymentMethod := ""
	if cus.InvoiceSettings.DefaultPaymentMethod != nil {
		paymentMethod = cus.InvoiceSettings.DefaultPaymentMethod.ID
	} else {
		cards, err := PaymentMethodsList(customerID)
		if err != nil {
			return "", err
		}
		if len(cards) == 0 {
			return "", errors.New("you must add a credit card to your billing account")
		}
		paymentMethod = cards[0].ID
	}

	return paymentMethod, nil
}

func CreatePaymentIntent(customerID, description string, amount int64) (string, error) {

	pm, err := GetDefaultPaymentMethod(customerID)
	if err != nil {
		return "", err
	}

	pi, err := paymentintent.New(&stripe.PaymentIntentParams{
		Amount:        stripe.Int64(amount),
		Customer:      stripe.String(customerID),
		PaymentMethod: stripe.String(pm),
		Currency:      stripe.String(string(stripe.CurrencyUSD)),
		Confirm:       stripe.Bool(false),
		Description:   stripe.String(description),
		CaptureMethod: stripe.String(string(stripe.PaymentIntentCaptureMethodManual)),
	})

	if err != nil {
		return "", err
	}
	return pi.ID, err
}

func ConfirmPaymentIntent(paymentintentID, customerID string) error {

	pm, err := GetDefaultPaymentMethod(customerID)
	if err != nil {
		return err
	}

	params := &stripe.PaymentIntentConfirmParams{
		PaymentMethod: stripe.String(pm),
	}

	_, err = paymentintent.Confirm(paymentintentID, params)
	if err != nil {
		return err
	}

	return nil
}

func CapturePaymentIntent(paymentintentID string, amount int64) error {

	params := &stripe.PaymentIntentCaptureParams{
		AmountToCapture: stripe.Int64(amount),
	}

	_, err := paymentintent.Capture(paymentintentID, params)
	if err != nil {
		return err
	}

	return nil
}

func CancelPaymentIntent(paymentintentID string, cancelReason string) error {

	params := &stripe.PaymentIntentCancelParams{
		CancellationReason: stripe.String(cancelReason),
	}

	_, err := paymentintent.Cancel(paymentintentID, params)
	if err != nil {
		return err
	}

	return nil
}

// ─────────────────────────────────────────────────────────────────────────────
func CreatePaymentMethod(customerID string, number, cvc string, expireMonth, expireYear int64) (string, error) {

	stripePrams := stripe.PaymentMethodParams{
		Card: &stripe.PaymentMethodCardParams{
			Number:   stripe.String(number),
			ExpMonth: stripe.Int64(expireMonth),
			ExpYear:  stripe.Int64(expireYear),
			CVC:      stripe.String(cvc),
		},
		Type: stripe.String(string(stripe.PaymentMethodTypeCard)),
	}

	pm, err := paymentmethod.New(&stripePrams)
	if err != nil {
		return "", err
	}

	_, err = paymentmethod.Attach(
		pm.ID,
		&stripe.PaymentMethodAttachParams{
			Customer: &customerID,
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

func PaymentIntentWithConfirmRefundFlow(sampleCustomerID string) error {

	//create payment intent
	pi, err := CreatePaymentIntent(sampleCustomerID, "pitest", 300)
	if err != nil {
		return err
	}

	//confirm for release
	err = ConfirmPaymentIntent(pi, sampleCustomerID)
	if err != nil {
		return err
	}

	//capture payment intent
	err = CapturePaymentIntent(pi, 300)
	if err != nil {
		return err
	}

	return nil

}

func PaymentIntentWithCancellFlow(sampleCustomerID string) error {

	//create payment intent
	pi, err := CreatePaymentIntent(sampleCustomerID, "pitest", 300)
	if err != nil {
		return err
	}

	//cancel payment
	reasone := string(stripe.PaymentIntentCancellationReasonRequestedByCustomer)
	err = CancelPaymentIntent(pi, reasone)
	if err != nil {
		return err
	}

	return nil
}
func main() {
	stripe.Key = "sk_test_sample"
	sampleCustomerID := "cus_sample"

	err := PaymentIntentWithConfirmRefundFlow(sampleCustomerID)
	if err != nil {
		panic(err)
	}

	err = PaymentIntentWithCancellFlow(sampleCustomerID)
	if err != nil {
		panic(err)
	}

}
