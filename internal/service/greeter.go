package service

import (
	"context"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	v1 "kratos-proto/api/helloworld/v1"
	"kratos-proto/internal/biz"
)

// GreeterService is a greeter service.
type GreeterService struct {
	v1.UnimplementedGreeterServer

	uc  *biz.GreeterUsecase
	log *log.Helper
}

// NewGreeterService new a greeter service.
func NewGreeterService(uc *biz.GreeterUsecase, logger log.Logger) *GreeterService {
	return &GreeterService{uc: uc, log: log.NewHelper(logger)}
}

// SayHello implements helloworld.GreeterServer
func (s *GreeterService) SayHello(ctx context.Context, in *v1.HelloRequest) (*v1.HelloReply, error) {
	s.log.Infof("SayHello Received: %v", in.GetName())
	if in.GetName() == "error" {
		return nil, errors.NotFound(v1.ErrorReason_USER_NOT_FOUND.String(), in.GetName())
	}
	return &v1.HelloReply{Message: "Hello " + in.GetName()}, nil
}

func (s *GreeterService) SayEmpty(ctx context.Context, _ *emptypb.Empty)  (*v1.HelloReply, error) {
	return &v1.HelloReply{Message: "Hello empty"}, nil
}

func (s *GreeterService) IsDefault(ctx context.Context, in *v1.HelloIsDefaultRequest) (*v1.HelloDefaultReply, error)  {
	var age int32
	if in.GetAge() != nil {
		age = in.GetAge().GetValue()
	}
	return &v1.HelloDefaultReply{
		Name: in.GetName(),
		Age:  age,
	}, nil
}

func (s *GreeterService) FiledMask(ctx context.Context, in *v1.HelloFieldMaskRequest) (*v1.HelloFieldMaskResponse, error)  {
	s.log.Infof("task_id = %+v", in.GetTaskId())
	s.log.Infof("is_finished = %+v", in.GetIsFinished())
	s.log.Infof("is_deleted = %+v", in.GetIsDelete())
	return &v1.HelloFieldMaskResponse{
		FieldMask: in.GetFieldMask().Paths,
	}, nil
}

func (s *GreeterService) AnyTypes(ctx context.Context, in *v1.HelloAnyTypesRequest) (*v1.HelloAnyTypesResponse, error)  {
	s.log.Infof("topic = %+v", in.GetTopic())
	s.log.Infof("desc = %s", in.GetDesc().GetValue())
	return &v1.HelloAnyTypesResponse{
		Topic: in.GetTopic(),
		Desc: string(in.GetDesc().GetValue()),
	}, nil
}

func (s *GreeterService) Times(ctx context.Context, in *v1.HelloTsRequest) (*v1.HelloTsResponse, error)  {
	s.log.Infof("seconds = %+v", in.GetTimeBegin().GetSeconds())
	s.log.Infof("nano = %s", in.GetTimeBegin().GetNanos())
	return &v1.HelloTsResponse{Timestamp: in.GetTimeBegin().GetSeconds()},nil
}