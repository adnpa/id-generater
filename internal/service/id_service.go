package service

import (
	"adnpa/id-generater/api/pb"
	"adnpa/id-generater/internal/global"
	"context"
)

type IdService struct {
	pb.UnimplementedIdServer
}

func (s *IdService) GetId(ctx context.Context, in *pb.GetIdReq) (*pb.GetIdResp, error) {
	id, err := global.Snowflake.Generate()
	if err != nil {
		panic(err)
	}
	return &pb.GetIdResp{Id: id}, nil
}
