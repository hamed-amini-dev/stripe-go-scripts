package stripe

import "github.com/hamed-amini-dev/stripe-go-scripts/internal/request"

// define sign interface functions
type IStripe interface {
	//Create Customer
	CreateCustomer(r *request.GenericRequest) (interface{}, error)
	//Add Payment Method to stripe
	AddPaymentMethod(r *request.GenericRequest) (interface{}, error)
	//create subscription to stripe
	CreateSubscription(r *request.GenericRequest) (interface{}, error)

	// PaymentIntentWithConfirmRefundFlow
	PaymentIntentWithConfirmRefundFlow(r *request.GenericRequest) (interface{}, error)

	// PaymentIntentWithCancelFlow
	PaymentIntentWithCancelFlow(r *request.GenericRequest) (interface{}, error)
}
