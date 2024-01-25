package controller

import (
	"context"

	"github.com/erupshis/golang-integration-developer-test/internal/service/client"
	"github.com/erupshis/golang-integration-developer-test/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Controller struct {
	pb.UnimplementedServiceServer

	client client.BaseClient
}

func NewController(client client.BaseClient) *Controller {
	return &Controller{
		client: client,
	}
}

func (c *Controller) GetBalance(ctx context.Context, in *pb.GetBalanceRequest) (*pb.GetBalanceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBalance not implemented")
}

func (c *Controller) SendBet(ctx context.Context, in *pb.SendBetRequest) (*pb.SendBetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendBet not implemented")
}
