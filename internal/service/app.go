package service

import (
	"context"
	pb "game/api/app/v1"
	"game/internal/biz"
)

type AppService struct {
	pb.UnimplementedAppServer
	ac *biz.AppUsecase
}

func NewAppService(ac *biz.AppUsecase) *AppService {
	return &AppService{
		ac: ac,
	}
}

func (s *AppService) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserReply, error) {
	return &pb.CreateUserReply{}, nil
}
