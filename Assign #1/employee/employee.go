package employee

type Employee interface {
	GetPosition() string
	SetPosition(position string)
	GetSalary() float64
	SetSalary(salary float64)
	GetAddress() string
	SetAddress(address string)
}

//
