package health

import (
	pb "github.com/butterv/kubernetes-sample1-app/app/gen/go/v1/health"
)

type healthService struct{}

// NewHealthService generates the `HealthServer` implementation.
func NewHealthService() pb.HealthServer {
	return &healthService{}
}
