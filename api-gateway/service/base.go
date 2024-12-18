package service

type Service struct {
	Auth     *AuthService
	Employee *EmployeeService
}

func New(auth *AuthService, employee *EmployeeService) *Service {
	return &Service{
		Auth:     auth,
		Employee: employee,
	}
}
