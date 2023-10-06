package softwareengineer

import "Go/employee"

type SoftwareEngineer struct {
	Position string
	Salary   float64
	Address  string
}

func NewEmployee3(position string, salary float64, address string) employee.Employee {
	return &SoftwareEngineer{position, salary, address}
}
func (s *SoftwareEngineer) GetPosition() string {
	return s.Position
}

func (s *SoftwareEngineer) SetPosition(position string) {
	s.Position = position
}

func (s *SoftwareEngineer) GetSalary() float64 {
	return s.Salary
}

func (s *SoftwareEngineer) SetSalary(salary float64) {
	s.Salary = salary
}

func (s *SoftwareEngineer) GetAddress() string {
	return s.Address
}

func (s *SoftwareEngineer) SetAddress(address string) {
	s.Address = address
}
