package stripe

import "github.com/stripe/stripe-go/v74"

type IStripe interface {
	//Create customer to stripe
	CreateCustomer(params *CreateCustomerParams) (string, error)
	//Add Payment Method to stripe
	AddPaymentMethod(params *AddPaymentMethodParams) (string, error)

	//create subscription to stripe
	CreateSubscription(params *CreateSubscriptionParams) (*stripe.Subscription, error)

	// PaymentIntentWithConfirmRefundFlow
	PaymentIntentWithConfirmRefundFlow(params *PaymentIntentWithConfirmRefundFlowParams) error

	// PaymentIntentWithCancellFlow
	PaymentIntentWithCancellFlow(params *PaymentIntentWithCancellFlowParams) error
}
