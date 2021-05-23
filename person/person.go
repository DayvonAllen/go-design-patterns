package person

type Human interface {
	SetName(string)
	GetName() string
	SetAge(int)
	GetAge() int
	SetBirthday(string)
	GetBirthday() string
}

type DefaultHumanImplementation struct {
	*Person
}

type Person struct {
	Name string
	Age int
	BirthDay string
}

func (p *Person) SetName(name string) {
	p.Name = name
}

func (p *Person) GetName() string {
	return p.Name
}

func (p *Person) SetAge(age int) {
	p.Age = age
}

func (p *Person) GetAge() int {
	return p.Age
}

func (p *Person) SetBirthday(birthday string) {
	p.BirthDay = birthday
}

func (p *Person) GetBirthday() string {
	return p.BirthDay
}

func CreateHuman(p *Person) *DefaultHumanImplementation {
	return &DefaultHumanImplementation{p}
}