package main

type Employee interface {
	GetPosition() string
	SetPosition(position string)
	GetSalary() float64
	SetSalary(salary float64)
	GetAddress() string
	SetAddress(address string)
}

type CEO struct {
	position string
	salary   float64
	address  string
}

func (c *CEO) GetPosition() string {
	return c.position
}

func (c *CEO) SetPosition(position string) {
	c.position = position
}

func (c *CEO) GetSalary() float64 {
	return c.salary
}

func (c *CEO) SetSalary(salary float64) {
	c.salary = salary
}

func (c *CEO) GetAddress() string {
	return c.address
}

func (c *CEO) SetAddress(address string) {
	c.address = address
}

type VP struct { //Vice president
	position string
	salary   float64
	address  string
}

func (v *VP) GetPosition() string {
	return v.position
}

func (v *VP) SetPosition(position string) {
	v.position = position
}

func (v *VP) GetSalary() float64 {
	return v.salary
}

func (v *VP) SetSalary(salary float64) {
	v.salary = salary
}

func (v *VP) GetAddress() string {
	return v.address
}

func (v *VP) SetAddress(address string) {
	v.address = address
}

type SoftwareEngineer struct {
	position string
	salary   float64
	address  string
}

func (s *SoftwareEngineer) GetPosition() string {
	return s.position
}

func (s *SoftwareEngineer) SetPosition(position string) {
	s.position = position
}

func (s *SoftwareEngineer) GetSalary() float64 {
	return s.salary
}

func (s *SoftwareEngineer) SetSalary(salary float64) {
	s.salary = salary
}

func (s *SoftwareEngineer) GetAddress() string {
	return s.address
}

func (s *SoftwareEngineer) SetAddress(address string) {
	s.address = address
}

type DataAnalyst struct {
	position string
	salary   float64
	address  string
}

func (r *DataAnalyst) GetPosition() string {
	return r.position
}

func (r *DataAnalyst) SetPosition(position string) {
	r.position = position
}

func (r *DataAnalyst) GetSalary() float64 {
	return r.salary
}

func (r *DataAnalyst) SetSalary(salary float64) {
	r.salary = salary
}

func (r *DataAnalyst) GetAddress() string {
	return r.address
}

func (r *DataAnalyst) SetAddress(address string) {
	r.address = address
}

type Manager struct {
	position  string
	salary    float64
	address   string
	employees []Employee
}

func (m *Manager) GetPosition() string {
	return m.position
}

func (m *Manager) SetPosition(position string) {
	m.position = position
}

func (m *Manager) GetSalary() float64 {
	return m.salary
}

func (m *Manager) SetSalary(salary float64) {
	m.salary = salary
}

func (m *Manager) GetAddress() string {
	return m.address
}

func (m *Manager) SetAddress(address string) {
	m.address = address
}

func (m *Manager) GetEmployees() []Employee {
	return m.employees
}
