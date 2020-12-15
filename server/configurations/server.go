package configurations

import (
	"context"

	"github.com/stephenzsy/doc-locker/server/common/app_context"
	configs "github.com/stephenzsy/doc-locker/server/common/configurations"
	configs_service "github.com/stephenzsy/doc-locker/server/gen/configurations"
	"google.golang.org/protobuf/types/known/structpb"
)

type server struct {
	configs_service.UnimplementedConfigurationsServiceServer
	serviceContext    app_context.AppContext
	deploymentConfigs configs.DeploymentConfigurationFile
}

func NewServer(ctx app_context.AppContext) (s server, err error) {
	data, err := configs.GetServerDeploymentConfigurationFile(ctx)
	s = server{
		serviceContext:    ctx,
		deploymentConfigs: data,
	}
	return
}

func (s *server) SiteConfigurations(context context.Context, req *configs_service.SiteConfigurationsRequest) (
	response *configs_service.SiteConfigurationsResponse, err error) {
	c, err := structpb.NewStruct(map[string](interface{}){
		"azure": map[string](interface{}){
			"applicationId": s.deploymentConfigs.Cloud.Azure.ApplicationId,
		},
	})
	if err != nil {
		return
	}
	response = &configs_service.SiteConfigurationsResponse{
		SiteConfigurations: c,
	}
	return
}
