package stripe

import (
	pStripe "github.com/hamed-amini-dev/stripe-go-scripts/pkg/stripe"
	"github.com/stripe/stripe-go/v74"
)

type iStripe struct {
}

var _ IStripe = &iStripe{}

func New(ops ...Option) (IStripe, error) {
	s := new(iStripe)
	for _, fn := range ops {
		err := fn(s)
		if err != nil {
			return nil, err
		}
	}
	return s, nil
}

func (s *iStripe) CreateCustomer(params *CreateCustomerParams) (string, error) {
	return pStripe.CreateCustomer(params.Name, params.Email)
}

func (s *iStripe) AddPaymentMethod(params *AddPaymentMethodParams) (string, error) {
	return pStripe.AddPaymentMethod(params.CustomerID, params.Number, params.Cvc, params.ExpMonth, params.ExpYear)
}

func (s *iStripe) CreateSubscription(params *CreateSubscriptionParams) (*stripe.Subscription, error) {
	return pStripe.CreateSubscription(params.PromotionCode, params.CustomerID, params.Price, params.TrialPeriodDays, params.Qty, params.Payout)
}

func (s *iStripe) PaymentIntentWithConfirmRefundFlow(params *PaymentIntentWithConfirmRefundFlowParams) error {
	return pStripe.PaymentIntentWithConfirmRefundFlow(params.CustomerID, params.Description, params.Amount)
}

func (s *iStripe) PaymentIntentWithCancellFlow(params *PaymentIntentWithCancellFlowParams) error {
	return pStripe.PaymentIntentWithCancellFlow(params.CustomerID, params.Description, params.Amount)
}
