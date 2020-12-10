package app_context

import (
	"context"
	"fmt"
	"os"
)

type appContextKey string

const (
	appContextKeyApp appContextKey = "app"
)

type appContextValue struct {
	elevated   bool
	caller     appContextCaller
	deployment appContextDeployment
}

const (
	WellKnownCallerdNone      string = "system:none"
	WellKnownCallerdAnonymous string = "system:anonymouse"
	WellKnownCallerdBootstrap string = "system:bootstrap"
)

type appContextCaller struct {
	id string
}

func (c appContextCaller) Id() string {
	return c.id
}

type appContextDeployment struct {
	id string
}

func (c appContextDeployment) Id() string {
	return c.id
}

type AppContext interface {
	context.Context
	Caller() appContextCaller
	Deployment() appContextDeployment
	Elevate() AppContext
	IsElevated() bool
}

type appContext struct {
	context.Context
}

func (ctx appContext) appContextValue() appContextValue {
	return ctx.Value(appContextKeyApp).(appContextValue)
}

func (ctx appContext) Caller() appContextCaller {
	return ctx.appContextValue().caller
}

func (ctx appContext) Deployment() appContextDeployment {
	return ctx.appContextValue().deployment
}

func (ctx appContext) Elevate() AppContext {
	nextContextValue := ctx.appContextValue()
	nextContextValue.elevated = true
	elevated := context.WithValue(ctx, appContextKeyApp, nextContextValue)
	return appContext{
		elevated,
	}
}

func (ctx appContext) IsElevated() bool {
	return ctx.appContextValue().elevated
}

func NewInitializeAppServiceContext(parent context.Context, callerId string, deploymentId string) AppContext {
	ctx := context.WithValue(parent, appContextKeyApp, appContextValue{
		elevated: false,
		caller: appContextCaller{
			id: callerId,
		},
		deployment: appContextDeployment{
			id: deploymentId,
		},
	})
	return appContext{
		ctx,
	}
}

func NewAppServiceContext(parent context.Context, callerId string) (AppContext, error) {
	deploymentId := os.Getenv("DOCLOCKER_DEPLOYMENT_ID")
	if deploymentId == "" {
		return nil, fmt.Errorf("missing environment variable: ")
	}

	return NewInitializeAppServiceContext(parent, callerId, deploymentId), nil
}
