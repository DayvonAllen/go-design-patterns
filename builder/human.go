package builder

// build aggregate example
type person struct {
	StreetAddress, Postcode, City string
	CompanyName, Position string
	Salary int
}

type PersonBuilder struct {
	*person
}

type PersonAddressBuilder struct {
	*PersonBuilder
}

type PersonJobBuilder struct {
	*PersonBuilder
}

func (b *PersonBuilder) Lives() *PersonAddressBuilder {
	return &PersonAddressBuilder{b}
}

func (b *PersonBuilder) Works() *PersonJobBuilder {
	return &PersonJobBuilder{b}
}

func (b *PersonAddressBuilder) At(streetAddress string) *PersonAddressBuilder {
	b.person.StreetAddress = streetAddress
	return b
}

func (b *PersonAddressBuilder) In(city string) *PersonAddressBuilder {
	b.person.City = city
	return b
}

func (b *PersonAddressBuilder) WithPostCode(postalCode string) *PersonAddressBuilder {
	b.person.Postcode = postalCode
	return b
}

func (it *PersonJobBuilder) For(company string) *PersonJobBuilder {
	it.person.CompanyName = company
	return it
}

func (it *PersonJobBuilder) As(position string) *PersonJobBuilder {
	it.person.Position = position
	return it
}

func (it *PersonJobBuilder) Makes(salary int) *PersonJobBuilder {
	it.person.Salary = salary
	return it
}

func (b *PersonBuilder) Build() *person {
	return b.person
}

// NewPersonBuilder factory function
func NewPersonBuilder() *PersonBuilder {
	return &PersonBuilder{&person{}}
}


