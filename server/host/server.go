package host

import (
	"context"
	"encoding/json"
	"runtime"

	pb "github.com/stephenzsy/doc-locker/server/gen/host"
)

type StatusInfo struct {
	GolangVersion string `json:"golangVersion"`
}

type HostServiceServer struct {
	pb.UnimplementedHostServiceServer
}

func (*HostServiceServer) Status(ctx context.Context, input *pb.HostStatusRequest) (*pb.HostStatusResponse, error) {
	statusJson, err := json.Marshal(StatusInfo{
		GolangVersion: runtime.Version(),
	})
	if err != nil {
		return nil, err
	}
	// No feature was found, return an unnamed feature
	return &pb.HostStatusResponse{StatusJson: string(statusJson)}, nil
}
