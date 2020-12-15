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

func VerifyCallerId(ctx AppContext, expectedCallerIds ...string) error {
	for _, expectedCallerId := range expectedCallerIds {
		if ctx.Caller().Id() == expectedCallerId {
			return nil
		}
	}
	return getVerificationError(fmt.Sprintf("callerId not allowed: %s", ctx.Caller().Id()))
}
