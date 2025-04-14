package test

import (
	"adnpa/id-generater/api/pb"
	"adnpa/id-generater/internal/utils"
	"context"
	"fmt"
	"testing"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func TestGeterate(t *testing.T) {
	snf, err := utils.NewSnowflake(1, 1744621641309)
	if err != nil {
		t.Error(err)
	}
	id, _ := snf.Generate()
	fmt.Printf("%b\n", id)
	t.Log(id)
	id, _ = snf.Generate()
	fmt.Printf("%b\n", id)
	t.Log(id)
}

func TestE2E(t *testing.T) {
	cc, err := grpc.NewClient("localhost:10000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Log("ee")
		t.Error(err)
	}
	cli := pb.NewIdClient(cc)
	pbid, err := cli.GetId(context.Background(), &pb.GetIdReq{})
	t.Log(pbid)
	t.Log(err)
}
