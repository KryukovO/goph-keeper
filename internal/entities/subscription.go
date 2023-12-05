package entities

import api "github.com/KryukovO/goph-keeper/api/serverpb"

// Subscription - подписка пользователя в файловом хранилище.
type Subscription string

const (
	// UnknownSubscription - неизвестная разновидность подписки.
	UnknownSubscription Subscription = "UNKNOWN"
	// RegularSubscription - обычная подписка.
	RegularSubscription Subscription = "REGULAR"
	// PremiumSubscription - премиум подписка.
	PremiumSubscription Subscription = "PREMIUM"
)

// MakeSubscription возвращает значение подписки, соответствуеющее значению из gRPC API.
func MakeSubscription(subscription api.Subscription) Subscription {
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

// ConvertSubscription конвертирует значение подписки в значение gRPC API.
func ConvertSubscription(subscription Subscription) api.Subscription {
	mapping := map[Subscription]api.Subscription{
		UnknownSubscription: api.Subscription_UNKNOWN,
		RegularSubscription: api.Subscription_REGULAR,
		PremiumSubscription: api.Subscription_PREMIUM,
	}

	if subscr, ok := mapping[subscription]; ok {
		return subscr
	}

	return api.Subscription_UNKNOWN
}
