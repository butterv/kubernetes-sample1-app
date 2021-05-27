package health_test

import (
	"context"
	"testing"

	"github.com/google/go-cmp/cmp"
	"google.golang.org/protobuf/testing/protocmp"

	"github.com/butterv/kubernetes-sample1-app/app/domain/service/health"
	pb "github.com/butterv/kubernetes-sample1-app/app/gen/go/v1/health"
)

func TestHealthService_Check(t *testing.T) {
	want := &pb.HealthCheckResponse{
		Status: pb.HealthCheckResponse_SERVING,
	}

	service := health.NewHealthService()

	ctx := context.Background()
	req := &pb.HealthCheckRequest{}

	got, err := service.Check(ctx, req)
	if err != nil {
		t.Fatalf("service.Check(ctx, %v) = _, %#v; want nil", req, err)
	}
	if diff := cmp.Diff(got, want, protocmp.Transform()); diff != "" {
		t.Errorf("service.Check(ctx, %v) = %#v, _; want %v\ndiff = %s", req, got, want, diff)
	}
}
