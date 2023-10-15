package manager

import "Go/employee"

type Manager struct {
	Position string
	Salary   float64
	Address  string
	//employees []employee.Employee
}

func NewEmployee2(position string, salary float64, address string) employee.Employee {
	return &Manager{position, salary, address}
}
func (m *Manager) GetPosition() string {
	return m.Position
}

func (m *Manager) SetPosition(position string) {
	m.Position = position
}

func (m *Manager) GetSalary() float64 {
	return m.Salary
}

func (m *Manager) SetSalary(salary float64) {
	m.Salary = salary
}

func (m *Manager) GetAddress() string {
	return m.Address
}

func (m *Manager) SetAddress(address string) {
	m.Address = address
}
