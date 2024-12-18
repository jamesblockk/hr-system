package client

import (
	"context"
	pb "hr-system/common/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type EmployeeServiceClient struct {
	client pb.EmployeeServiceClient
	conn   *grpc.ClientConn
}

func NewEmployeeServiceClient(address string) (*EmployeeServiceClient, error) {
	conn, err := grpc.NewClient(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	return &EmployeeServiceClient{
		client: pb.NewEmployeeServiceClient(conn),
		conn:   conn,
	}, nil
}

func (c *EmployeeServiceClient) Close() error {
	return c.conn.Close()
}

func (c *EmployeeServiceClient) GetEmployee(ctx context.Context, req *pb.GetEmployeeRequest) (*pb.EmployeeResponse, error) {
	return c.client.GetEmployeeByID(ctx, req)
}
