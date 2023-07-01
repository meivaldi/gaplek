package grpc

import (
	"context"
	"errors"

	service "github.com/meivaldi/gaplek/internal/service"
	pb "github.com/meivaldi/protobuf/gaplek"
)

type JitterHandler struct {
	pb.UnimplementedJitterServer
	services *service.Service
}

func NewJitterHandler(svc *service.Service) pb.JitterServer {
	return &JitterHandler{
		services: svc,
	}
}

func (handler *JitterHandler) GetJitter(ctx context.Context, req *pb.JitterRequest) (res *pb.JitterResponse, err error) {
	if req == nil {
		return nil, errors.New("empty payload")
	}

	if req.High <= 0 {
		return nil, errors.New("range must be not zero number")
	}

	jitter := handler.services.JitterService.GetRandomNumber(req.Low, req.High)

	return &pb.JitterResponse{
		Jitter: jitter,
	}, nil
}
