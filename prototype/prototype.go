package prototype

import (
	"bytes"
	"encoding/gob"
)

type Address struct {
	StreetAddress, City, Country string
}

func (a *Address) NaiveDeepCopy() *Address {
	return &Address{
		a.StreetAddress,
		a.City,
		a.Country }
}

type Person struct {
	Name string
	Address *Address
	Friends []string
}

func (p *Person) NaiveDeepCopy() *Person {
	c := *p
	c.Address = p.Address.NaiveDeepCopy()
	// copy the list of friends from the original slice and gives "c" a copy of that original slice and creates
	// a new object in memory
	copy(c.Friends, p.Friends)
	return &c
}

// BetterDeepCopy Serialization approach
func (p *Person) BetterDeepCopy() *Person {
	// need to work with bytes for serialization
	b := bytes.Buffer{}

	// creating the encoder, bytes will be stored in b
	e := gob.NewEncoder(&b)

	// encodes(serializes the person object into a gob
	_ = e.Encode(p)

	// creating the decode
	d := gob.NewDecoder(&b)

	// prepare memory for the person object that will be created from the decoder
	result := new(Person)

	// decoders need pointers to objects or structs
	_ = d.Decode(result)

	return result
}