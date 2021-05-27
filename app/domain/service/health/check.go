package health

import (
	"context"

	pb "github.com/butterv/kubernetes-sample1-app/app/gen/go/v1/health"
)

func (*healthService) Check(context.Context, *pb.HealthCheckRequest) (*pb.HealthCheckResponse, error) {
	// 	0: UNKNOWN
	//	1: SERVING
	//	2: NOT_SERVING
	//	3: SERVICE_UNKNOWN
	return &pb.HealthCheckResponse{
		Status: pb.HealthCheckResponse_SERVING,
	}, nil
}
