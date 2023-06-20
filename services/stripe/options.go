package stripe

import (
	pStripe "github.com/hamed-amini-dev/stripe-go-scripts/pkg/stripe"
)

type Option func(istripe *iStripe) error

// InitOptionAccountModel for initialize model account business logic

func InitOptionStripeKey(key string) Option {
	return func(istripe *iStripe) error {
		pStripe.SetKey(key)
		return nil
	}
}
