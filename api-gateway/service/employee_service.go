package service

import (
	"context"
	"hr-system/api-gateway/client"
	pb "hr-system/common/proto"
)

type EmployeeService struct {
	employeeClient *client.EmployeeServiceClient
}

func NewEmployeeService(employeeClient *client.EmployeeServiceClient) *EmployeeService {
	return &EmployeeService{employeeClient: employeeClient}
}

func (s *EmployeeService) GetEmployeeByID(ctx context.Context, id uint64) (*pb.Employee, error) {
	req := &pb.GetEmployeeRequest{Id: id}
	resp, err := s.employeeClient.GetEmployee(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.Employee, nil
}
