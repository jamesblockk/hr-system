package internal_test

import (
	"context"
	"hr-system/common/config"
	"hr-system/common/dao/query"
	"hr-system/common/database"
	"hr-system/common/proto"
	"hr-system/employee-service/internal"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockEmployeeQuery struct {
	mock.Mock
}

func (m *MockEmployeeQuery) WithContext(ctx context.Context) *MockEmployeeQuery {
	return m
}

func (m *MockEmployeeQuery) Where(condition interface{}) *MockEmployeeQuery {
	return m
}

func (m *MockEmployeeQuery) First() (*Employee, error) {
	args := m.Called()
	return args.Get(0).(*Employee), args.Error(1)
}

// Mock Employee struct
type Employee struct {
	ID           uint
	Name         string
	Email        string
	Phone        string
	DepartmentID uint
	PositionID   uint
	HireDate     string // Modify this if you have a more specific type for date
	Salary       float64
	Department   *Department
	Position     *Position
}

type Department struct {
	ID   uint
	Name string
}

type Position struct {
	ID   uint
	Name string
}

func TestGetEmployeeByID(t *testing.T) {
	database.Init(config.GetLocal())
	query.SetDefault(database.GetDB())

	// Create mock data
	mockEmployee := &Employee{
		ID:           1,
		Name:         "Alice",
		Email:        "alice@example.com",
		Phone:        "111222333",
		DepartmentID: 1,
		PositionID:   1,
		HireDate:     "2021-01-01", // Mocked date
		Salary:       50000,
		Department:   &Department{ID: 1, Name: "Engineering"},
		Position:     &Position{ID: 1, Name: "Developer"},
	}

	// Create a mock EmployeeQuery
	mockQuery := new(MockEmployeeQuery)
	mockQuery.On("WithContext", mock.Anything).Return(mockQuery)
	mockQuery.On("Where", mock.Anything).Return(mockQuery)
	mockQuery.On("First").Return(mockEmployee, nil)

	server := &internal.Server{}

	// Mock request
	req := &proto.GetEmployeeRequest{Id: 1}

	// Call GetEmployeeByID method
	resp, err := server.GetEmployeeByID(context.Background(), req)

	// Assert the response
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, uint64(1), resp.Employee.Id)
	assert.Equal(t, "Alice", resp.Employee.Name)
	assert.Equal(t, "alice@example.com", resp.Employee.Email)
	assert.Equal(t, "111222333", resp.Employee.Phone)
	assert.Equal(t, uint64(1), resp.Employee.DepartmentId)
	assert.Equal(t, uint64(1), resp.Employee.PositionId)
	assert.Equal(t, "2021-01-01", resp.Employee.HireDate)
	assert.Equal(t, 50000.0, resp.Employee.Salary)
	assert.Equal(t, "Engineering", resp.Employee.Department.Name)
	assert.Equal(t, "Developer", resp.Employee.Position.Name)

	// Assert that the mock methods were called
	mockQuery.AssertExpectations(t)
}
