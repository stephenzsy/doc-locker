package app_context

import (
	"fmt"
)

func getVerificationError(reason string) error {
	return fmt.Errorf("Context verification failure: %s", reason)
}

func VerifyElevated(ctx AppContext) error {
	if !ctx.IsElevated() {
		return getVerificationError("not elevated")
	}
	return nil
}

func VerifyCallerId(ctx AppContext, expectedCallerId string) error {
	if ctx.Caller().Id() != expectedCallerId {
		return getVerificationError("callerId not allowed")
	}
	return nil
}
