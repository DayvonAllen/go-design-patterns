package observer

import (
	"container/list"
	"fmt"
)

type ObservableDep struct {
	subs *list.List
}

func (o *ObservableDep) Subscribe(x ObserverDep) {
	o.subs.PushBack(x)
}

func (o *ObservableDep) Unsubscribe(x ObserverDep) {
	for z := o.subs.Front(); z != nil; z = z.Next() {
		if z.Value.(ObserverDep) == x {
			o.subs.Remove(z)
		}
	}
}

func (o *ObservableDep) Fire(data interface{}) {
	for z := o.subs.Front(); z != nil; z = z.Next() {
		z.Value.(ObserverDep).Notify(data)
	}
}

type ObserverDep interface {
	Notify(data interface{})
}

type PersonDep struct {
	ObservableDep
	age int
}

func NewPersonDep(age int) *PersonDep {
	return &PersonDep{ObservableDep{new(list.List)}, age}
}

type PropertyChangedDep struct {
	Name  string
	Value interface{}
}

func (p *PersonDep) Age() int { return p.age }
func (p *PersonDep) SetAge(age int) {
	if age == p.age {
		return
	} // no change

	oldCanVote := p.CanVote()

	p.age = age
	p.Fire(PropertyChangedDep{"Age", p.age})

	if oldCanVote != p.CanVote() {
		p.Fire(PropertyChangedDep{"CanVote", p.CanVote()})
	}
}

func (p *PersonDep) CanVote() bool {
	return p.age >= 18
}

type ElectricalRoll struct {
}

func (e *ElectricalRoll) Notify(data interface{}) {
	if pc, ok := data.(PropertyChangedDep); ok {
		if pc.Name == "CanVote" && pc.Value.(bool) {
			fmt.Println("Congratulations, you can vote!")
		}
	}
}

func StartDependencyObserver() {
	p := NewPersonDep(0)
	er := &ElectricalRoll{}
	p.Subscribe(er)

	for i := 10; i < 20; i++ {
		fmt.Println("Setting age to", i)
		p.SetAge(i)
	}
}
