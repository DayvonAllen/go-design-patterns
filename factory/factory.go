package factory

type Employee struct {
	Name, Position string
	Salary int
}

// NewEmployeeFactory / functional
// this is a function that returns a function or in other words a higher order function
func NewEmployeeFactory(position string, salary int) func(name string) *Employee {
	return func(name string) *Employee {
		return &Employee{name, position, salary}
	}
}

// EmployeeFactory structured factory approach
type EmployeeFactory struct {
	Position string
	Salary int
}

func NewStructuredEmployeeFactory(position string, salary int)  *EmployeeFactory {
	return &EmployeeFactory{position, salary}
}


func (f *EmployeeFactory) Create(name string) *Employee {
	return &Employee{name, f.Position, f.Salary}
}