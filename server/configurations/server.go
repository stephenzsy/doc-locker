package configurations

import (
	"context"
	"encoding/json"

	"github.com/stephenzsy/doc-locker/server/common/app_context"
	configs "github.com/stephenzsy/doc-locker/server/common/configurations"
	configs_service "github.com/stephenzsy/doc-locker/server/gen/configurations"
	"google.golang.org/protobuf/types/known/emptypb"
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
	configs.SharedServerCloudAwsConfiguration
}

type siteConfigurationsData struct {
	Aws siteConfigurationsAwsData `json:"aws"`
}

func (s *server) SiteConfigurations(context context.Context, req *emptypb.Empty) (
	response *configs_service.SiteConfigurationsResponse, err error) {
	data := siteConfigurationsData{
		Aws: siteConfigurationsAwsData{
			SharedServerCloudAwsConfiguration: s.deploymentConfigs.Cloud.Aws.SharedServerCloudAwsConfiguration,
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
