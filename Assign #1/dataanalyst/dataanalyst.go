package dataanalyst

import "Go/employee"

type DataAnalyst struct {
	Position string
	Salary   float64
	Address  string
}

func NewEmployee1(position string, salary float64, address string) employee.Employee {
	return &DataAnalyst{position, salary, address}
}
func (d *DataAnalyst) GetPosition() string {
	return d.Position
}

func (d *DataAnalyst) SetPosition(position string) {
	d.Position = position
}

func (d *DataAnalyst) GetSalary() float64 {
	return d.Salary
}

func (d *DataAnalyst) SetSalary(salary float64) {
	d.Salary = salary
}

func (d *DataAnalyst) GetAddress() string {
	return d.Address
}

func (d *DataAnalyst) SetAddress(address string) {
	d.Address = address
}
