package internal

import (
	"context"
	"hr-system/common/dao/query"
	"hr-system/common/proto"

	"github.com/jinzhu/copier"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"

	"log"
)

type Server struct {
	proto.UnimplementedEmployeeServiceServer
}

func (s *Server) GetEmployeeByID(ctx context.Context, req *proto.GetEmployeeRequest) (result *proto.EmployeeResponse, err error) {
	log.Printf("GetEmployee called with: %v", req.GetId())

	employeeQuery := query.Employee
	employee, err := employeeQuery.WithContext(ctx).
		Preload(employeeQuery.Department).
		Preload(employeeQuery.Position).
		Where(
			employeeQuery.ID.Eq(uint(req.GetId())),
		).First()
	if err != nil {
		return
	}

	result = &proto.EmployeeResponse{Employee: &proto.Employee{}}
	if err = copier.CopyWithOption(&result.Employee, &employee, copier.Option{}); err != nil {
		return
	}

	result.Employee.HireDate = timestamppb.New(employee.HireDate)

	return result, nil
}
