package vp

import "Go/employee"

type VP struct { //Vice president
	Position string
	Salary   float64
	Address  string
}

func NewEmployee4(position string, salary float64, address string) employee.Employee {
	return &VP{position, salary, address}
}
func (v *VP) GetPosition() string {
	return v.Position
}

func (v *VP) SetPosition(position string) {
	v.Position = position
}

func (v *VP) GetSalary() float64 {
	return v.Salary
}

func (v *VP) SetSalary(salary float64) {
	v.Salary = salary
}

func (v *VP) GetAddress() string {
	return v.Address
}

func (v *VP) SetAddress(address string) {
	v.Address = address
}
