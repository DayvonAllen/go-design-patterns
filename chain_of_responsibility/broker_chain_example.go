package chain_of_responsibility

import (
	"fmt"
	"sync"
)

// cqs, mediator, cor

type Argument int

const (
	Attack Argument = iota
	Defense
)

type Query struct {
	CreatureName string
	WhatToQuery  Argument
	Value        int
}

type Observer interface {
	Handle(*Query)
}

type Observable interface {
	Subscribe(o Observer)
	Unsubscribe(o Observer)
	Fire(q *Query)
}

type Game struct {
	observers sync.Map
}

func (g *Game) Subscribe(o Observer) {
	g.observers.Store(o, struct{}{})
	//                   ↑↑↑ empty anon struct
}

func (g *Game) Unsubscribe(o Observer) {
	g.observers.Delete(o)
}

func (g *Game) Fire(q *Query) {
	g.observers.Range(func(key, value interface{}) bool {
		if key == nil {
			return false
		}
		key.(Observer).Handle(q)
		return true
	})
}

type CreatureBroker struct {
	game            *Game
	Name            string
	attack, defense int // ← private!
}

func NewCreatureBroker(game *Game, name string, attack int, defense int) *CreatureBroker {
	return &CreatureBroker{game: game, Name: name, attack: attack, defense: defense}
}

func (c *CreatureBroker) Attack() int {
	q := Query{c.Name, Attack, c.attack}
	c.game.Fire(&q)
	return q.Value
}

func (c *CreatureBroker) Defense() int {
	q := Query{c.Name, Defense, c.defense}
	c.game.Fire(&q)
	return q.Value
}

func (c *CreatureBroker) String() string {
	return fmt.Sprintf("%s (%d/%d)",
		c.Name, c.Attack(), c.Defense())
}

// data common to all modifiers
type CreatureModifierBroker struct {
	game     *Game
	creature *CreatureBroker
}

func (c *CreatureModifierBroker) Handle(*Query) {
	// nothing here!
}

type DoubleAttackModifierBroker struct {
	CreatureModifierBroker
}

func NewDoubleAttackModifierBroker(g *Game, c *CreatureBroker) *DoubleAttackModifierBroker {
	d := &DoubleAttackModifierBroker{CreatureModifierBroker{g, c}}
	g.Subscribe(d)
	return d
}

func (d *DoubleAttackModifierBroker) Handle(q *Query) {
	if q.CreatureName == d.creature.Name &&
		q.WhatToQuery == Attack {
		q.Value *= 2
	}
}

func (d *DoubleAttackModifierBroker) Close() error {
	d.game.Unsubscribe(d)
	return nil
}

func CreateBrokerChain() {
	game := &Game{sync.Map{}}
	goblin := NewCreatureBroker(game, "Strong Goblin", 2, 2)
	fmt.Println(goblin.String())

	{
		m := NewDoubleAttackModifierBroker(game, goblin)
		fmt.Println(goblin.String())
		m.Close()
	}

	fmt.Println(goblin.String())
}
