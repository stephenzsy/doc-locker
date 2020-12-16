package configurations

import (
	"context"
	"encoding/json"

	"github.com/stephenzsy/doc-locker/server/common/app_context"
	configs "github.com/stephenzsy/doc-locker/server/common/configurations"
	configs_service "github.com/stephenzsy/doc-locker/server/gen/configurations"
)

type server struct {
	configs_service.UnimplementedConfigurationsServiceServer
	serviceContext    app_context.AppContext
	deploymentConfigs configs.DeploymentConfigurationFile
}

func NewServer(ctx app_context.AppContext) (s server, err error) {
	data, err := configs.GetServerDeploymentConfigurationFile(ctx.Elevate())
	s = server{
		serviceContext:    ctx,
		deploymentConfigs: data,
	}
	return
}

type siteConfigurationsAwsData struct {
	CognitoIdentityPoolId      string `json:"cognitoIdentityPoolId"`
	CognitoRegion              string `json:"cognitoRegion"`
	CognitoUserPoolId          string `json:"cognitoUserPoolId"`
	CognitoUserPoolWebClientId string `json:"cognitoUserPoolWebClientId"`
}

type siteConfigurationsData struct {
	Aws siteConfigurationsAwsData `json:"aws"`
}

func (s *server) SiteConfigurations(context context.Context, req *configs_service.SiteConfigurationsRequest) (
	response *configs_service.SiteConfigurationsResponse, err error) {
	data := siteConfigurationsData{
		Aws: siteConfigurationsAwsData{
			CognitoIdentityPoolId:      s.deploymentConfigs.Cloud.Aws.CognitoIdentityPoolId,
			CognitoRegion:              s.deploymentConfigs.Cloud.Aws.CognitoRegion,
			CognitoUserPoolId:          s.deploymentConfigs.Cloud.Aws.CognitoUserPoolId,
			CognitoUserPoolWebClientId: s.deploymentConfigs.Cloud.Aws.CognitoUserPoolWebClientId,
		},
	}
	if err != nil {
		return
	}
	marshalled, err := json.Marshal(data)
	if err != nil {
		return
	}
	response = &configs_service.SiteConfigurationsResponse{
		SiteConfigurationsJson: string(marshalled),
	}
	return
}
