package longrunning

import (
	"context"
	"fmt"
	extLongrunning "google.golang.org/genproto/googleapis/longrunning"
	"google.golang.org/protobuf/types/known/emptypb"
)

type server struct{}

func (s *server) ListOperations(ctx context.Context, request *extLongrunning.ListOperationsRequest) (*extLongrunning.ListOperationsResponse, error) {
	fmt.Println("request in ListOperations")
	return nil, nil
}

func (s *server) GetOperation(ctx context.Context, request *extLongrunning.GetOperationRequest) (*extLongrunning.Operation, error) {
	fmt.Println("request in GetOperation")
	return nil, nil
}

func (s *server) DeleteOperation(ctx context.Context, request *extLongrunning.DeleteOperationRequest) (*emptypb.Empty, error) {
	//TODO implement me
	panic("implement me")
}

func (s *server) CancelOperation(ctx context.Context, request *extLongrunning.CancelOperationRequest) (*emptypb.Empty, error) {
	//TODO implement me
	panic("implement me")
}

func (s *server) WaitOperation(ctx context.Context, request *extLongrunning.WaitOperationRequest) (*extLongrunning.Operation, error) {
	//TODO implement me
	panic("implement me")
}
