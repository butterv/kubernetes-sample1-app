package status_test

import (
	"errors"
	"reflect"
	"testing"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	appstatus "github.com/butterv/kubernetes-sample1-app/app/domain/service/status"
)

func TestMust(t *testing.T) {
	want := status.New(codes.NotFound, "TEST_ERROR_MESSAGE")

	s := status.New(codes.NotFound, "TEST_ERROR_MESSAGE")
	var err error

	got := appstatus.ExportMust(s, err)
	if got == nil {
		t.Fatalf("appstatus.ExportMust(%v, %v) = nil; want %v", s, err, want)
	}
	if !errors.Is(got.Err(), want.Err()) {
		t.Errorf("got.Err() = %#v; want %v", got.Err(), want.Err())
	}
	if got.Code() != want.Code() {
		t.Errorf("got.Code() = %d; want %d", got.Code(), want.Code())
	}
	if got.Message() != want.Message() {
		t.Errorf("got.Message() = %s; want %s", got.Message(), want.Message())
	}
}

func TestMust_Panic(t *testing.T) {
	defer func() {
		wantErr := errors.New("an error occurred")

		p := recover()
		if p == nil {
			t.Fatalf("no panic is detected. panic should be called. want %v", wantErr)
		}
		if !reflect.DeepEqual(p, wantErr) {
			t.Errorf("unexpected error is returned; got %#v, want %v", p, wantErr)
		}
	}()

	var s *status.Status
	err := errors.New("an error occurred")
	_ = appstatus.ExportMust(s, err)
}
