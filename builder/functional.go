package builder

type Person struct {
	name, position string
}

type personMod func(person2 *Person)

type FunctionalPersonBuilder struct {
	actions []personMod
}

func (b *FunctionalPersonBuilder) Called(name string) *FunctionalPersonBuilder {
	b.actions = append(b.actions, func(p *Person) {
		p.name = name
	})
	return b
}

func (b *FunctionalPersonBuilder) As(position string) *FunctionalPersonBuilder {
	b.actions = append(b.actions, func(p *Person) {
		p.position = position
	})
	return b
}

func (b *FunctionalPersonBuilder) Build() *Person {
	p := new(Person)
	for _, a := range b.actions {
		a(p)
	}
	return p
}
