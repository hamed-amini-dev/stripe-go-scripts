package stripe

import (
	"github.com/stripe/stripe-go/v74"
)

func SetKey(key string) {
	stripe.Key = key
}
