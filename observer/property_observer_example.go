package observer

import (
	"container/list"
	"fmt"
)

type ObservableProperty struct {
	subs *list.List
}

func (o *ObservableProperty) Subscribe(x ObserverProperty) {
	o.subs.PushBack(x)
}

func (o *ObservableProperty) Unsubscribe(x ObserverProperty) {
	for z := o.subs.Front(); z != nil; z = z.Next() {
		if z.Value.(ObserverProperty) == x {
			o.subs.Remove(z)
		}
	}
}

func (o *ObservableProperty) Fire(data interface{}) {
	for z := o.subs.Front(); z != nil; z = z.Next() {
		z.Value.(ObserverProperty).Notify(data)
	}
}

type ObserverProperty interface {
	Notify(data interface{})
}

type PersonProperty struct {
	ObservableProperty
	age int
}

func NewPersonProperty(age int) *PersonProperty {
	return &PersonProperty{ObservableProperty{new(list.List)}, age}
}

type PropertyChanged struct {
	Name  string
	Value interface{}
}

func (p *PersonProperty) Age() int { return p.age }
func (p *PersonProperty) SetAge(age int) {
	if age == p.age {
		return
	} // no change
	p.age = age
	p.Fire(PropertyChanged{"Age", p.age})
}

type TrafficManagement struct {
	o ObservableProperty
}

func (t *TrafficManagement) Notify(data interface{}) {
	if pc, ok := data.(PropertyChanged); ok {
		if pc.Value.(int) >= 16 {
			fmt.Println("Congrats, you can drive now!")
			// we no longer care
			t.o.Unsubscribe(t)
		}
	}
}

func StartPropertyObserver() {
	p := NewPersonProperty(15)
	t := &TrafficManagement{p.ObservableProperty}
	p.Subscribe(t)

	for i := 16; i <= 20; i++ {
		fmt.Println("Setting age to", i)
		p.SetAge(i)
	}
}
