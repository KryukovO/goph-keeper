package entities

import api "github.com/KryukovO/goph-keeper/api/serverpb"

type Subscription string

const (
	UnknownSubscription Subscription = "UNKNOWN"
	RegularSubscription Subscription = "REGULAR"
	PremiumSubscription Subscription = "PREMIUM"
)

func ConvertSubscription(subscription api.Subscription) Subscription {
	mapping := map[api.Subscription]Subscription{
		api.Subscription_UNKNOWN: UnknownSubscription,
		api.Subscription_REGULAR: RegularSubscription,
		api.Subscription_PREMIUM: PremiumSubscription,
	}

	if subscr, ok := mapping[subscription]; ok {
		return subscr
	}

	return UnknownSubscription
}
