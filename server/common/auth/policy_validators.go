package auth

const (
	WellKnownCallerIdNone      string = "none"
	WellKnownCallerIdAnonymous string = "anonymous"
	SystemCallerIdBootstrap    string = "system:bootstrap"
	ServiceCallerIdSds         string = "service:sds"
)

type AuthorizationPolicyValidator interface {
	Validate(callerId string, resourceId string) error
}
