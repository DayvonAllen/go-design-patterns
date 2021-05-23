package prototype

import (
	"bytes"
	"encoding/gob"
)

// prototypes
var mainOffice = FactoryEmployee{"", &FactoryAddress{0,"123 1st St.", "Miami"}}
var branchOffice = FactoryEmployee{"", &FactoryAddress{0,"438 3rd St.", "Houston"}}

type FactoryAddress struct {
	Suite int
	StreetAddress, City string
}

type FactoryEmployee struct {
	Name string
	Office *FactoryAddress
}

func newEmployee(proto *FactoryEmployee, name string, suite int) *FactoryEmployee {
	result := proto.DeepCopy()
	result.Name = name
	result.Office.Suite = suite

	return  result
}

func NewMainOfficeEmployee(name string, suite int) *FactoryEmployee {
	return  newEmployee(&mainOffice, name, suite)
}

func NewBranchOfficeEmployee(name string, suite int) *FactoryEmployee {
	return  newEmployee(&branchOffice, name, suite)
}

func (em *FactoryEmployee) DeepCopy() *FactoryEmployee {
	// need to work with bytes for serialization
	b := bytes.Buffer{}

	// creating the encoder, bytes will be stored in b
	e := gob.NewEncoder(&b)

	// encodes(serializes the person object into a gob
	_ = e.Encode(em)

	// creating the decode
	d := gob.NewDecoder(&b)

	// prepare memory for the person object that will be created from the decoder
	result := new(FactoryEmployee)

	// decoders need pointers to objects or structs
	_ = d.Decode(result)

	return result
}