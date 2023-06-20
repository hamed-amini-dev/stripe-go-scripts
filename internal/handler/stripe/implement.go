package stripe

/*
define endpoint functions
we need service module for using business logic of endpoint
*/
import (
	"encoding/json"

	"github.com/hamed-amini-dev/stripe-go-scripts/internal/request"
	sStripe "github.com/hamed-amini-dev/stripe-go-scripts/services/stripe"
)

type iStripe struct {
	service sStripe.IStripe
}

// create object handler for handling route
// need option for getting service module
// return handler object for using logic functionality

func New(ops ...Option) (IStripe, error) {
	h := new(iStripe)
	for _, fn := range ops {
		err := fn(h)
		if err != nil {
			return nil, err
		}
	}
	return h, nil
}

func (is iStripe) CreateCustomer(r *request.GenericRequest) (interface{}, error) {
	var params sStripe.CreateCustomerParams
	err := json.Unmarshal(r.Body, &params)
	if err != nil {
		return nil, err
	}
	return is.service.CreateCustomer(&params)
}

func (is iStripe) AddPaymentMethod(r *request.GenericRequest) (interface{}, error) {
	var params sStripe.AddPaymentMethodParams
	err := json.Unmarshal(r.Body, &params)
	if err != nil {
		return nil, err
	}
	return is.service.AddPaymentMethod(&params)
}

func (is iStripe) CreateSubscription(r *request.GenericRequest) (interface{}, error) {
	var params sStripe.CreateSubscriptionParams
	err := json.Unmarshal(r.Body, &params)
	if err != nil {
		return nil, err
	}
	return is.service.CreateSubscription(&params)
}

func (is iStripe) PaymentIntentWithConfirmRefundFlow(r *request.GenericRequest) (interface{}, error) {
	var params sStripe.PaymentIntentWithConfirmRefundFlowParams
	err := json.Unmarshal(r.Body, &params)
	if err != nil {
		return nil, err
	}
	return nil, is.service.PaymentIntentWithConfirmRefundFlow(&params)
}

func (is iStripe) PaymentIntentWithCancellFlow(r *request.GenericRequest) (interface{}, error) {
	var params sStripe.PaymentIntentWithCancellFlowParams
	err := json.Unmarshal(r.Body, &params)
	if err != nil {
		return nil, err
	}
	return nil, is.service.PaymentIntentWithCancellFlow(&params)
}
