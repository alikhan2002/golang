package ceo

import "Go/employee"

type CEO struct {
	Position string
	Salary   float64
	Address  string
}

func NewEmployee5(position string, salary float64, address string) employee.Employee {
	return &CEO{position, salary, address}

}
func (c *CEO) GetPosition() string {
	return c.Position
}

func (c *CEO) SetPosition(position string) {
	c.Position = position
}

func (c *CEO) GetSalary() float64 {
	return c.Salary
}

func (c *CEO) SetSalary(salary float64) {
	c.Salary = salary
}

func (c *CEO) GetAddress() string {
	return c.Address
}

func (c *CEO) SetAddress(address string) {
	c.Address = address
}
