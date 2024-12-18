package client

import (
	"context"

	pb "hr-system/common/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type AuthServiceClient struct {
	client pb.AuthServiceClient
	conn   *grpc.ClientConn
}

func NewAuthServiceClient(address string) (*AuthServiceClient, error) {
	conn, err := grpc.NewClient(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	return &AuthServiceClient{
		client: pb.NewAuthServiceClient(conn),
		conn:   conn,
	}, nil
}

func (c *AuthServiceClient) Login(ctx context.Context, email, password string) (string, error) {
	req := &pb.LoginRequest{Email: email, Password: password}

	resp, err := c.client.Login(ctx, req)
	if err != nil {
		return "", err
	}
	return resp.Token, nil
}

func (c *AuthServiceClient) Register(ctx context.Context, email, password string) error {
	req := &pb.RegisterRequest{Email: email, Password: password}
	_, err := c.client.Register(ctx, req)
	return err
}

func (c *AuthServiceClient) Close() error {
	return c.conn.Close()
}
