package main

import (
	"com.example/app/adpater"
	"com.example/app/builder"
	"com.example/app/factory"
	"com.example/app/prototype"
	"fmt"
)

func main() {
	adapterExample()
}

func customHtmlBuilderExample() {
	b := builder.NewHtmlBuilder("ul")

	//fluent way, chaining calls
	// reuse the result of a method call
	b.AddChildFluent("li", "hello").
		AddChildFluent("li", "world").
		AddChildFluent("li", "car").
		AddChildFluent("li", "bus").
		AddChildFluent("li", "yellow")
	fmt.Println(b.String())
}

// more complex builder example, uses 3 builders
func customPersonBuilderExample() {
	pb := builder.NewPersonBuilder()
	pb.
		Lives().
			At("145 S. St.").
			In("Houston").
			WithPostCode("34627").
		Works().
			For("Google").
			As("Programmer").
			Makes(120000)

	newPerson := pb.Build()

	fmt.Println(newPerson)
}

func functionalBuilder() {
	fb := builder.FunctionalPersonBuilder{}
	p := fb.
		Called("Jake").
		As("Programmer").
		Build()

	fmt.Println(p)
}

func customFunctionalFactory() {
	developerFactory := factory.NewEmployeeFactory("developer", 60000)
	managerFactory := factory.NewEmployeeFactory("manager", 80000)

	developer := developerFactory("John Doe")
	manager := managerFactory("Beth Johnson")

	fmt.Println(developer)
	fmt.Println(manager)
}

func customStructuredFactory() {
	developerFactory := factory.NewStructuredEmployeeFactory("developer", 60000)
	managerFactory := factory.NewStructuredEmployeeFactory("manager", 80000)

	developer := developerFactory.Create("Jerry")

	manager := managerFactory.Create("Jessica")

	fmt.Println(developer)
	fmt.Println(manager)
}

// prototypes require the ability to make deep copy of objects, types and structs
func naivePrototypeDeepCopyImpl() {
	john := prototype.Person{Name: "John",
		Address:     &prototype.Address{StreetAddress: "321 1st St.", City: "Houston", Country: "USA"},
			Friends: []string{"Chris", "Matt"},
		}

		// gives jane a copy of everything in john
		jane := john.NaiveDeepCopy()
		jane.Name = "Jane"
		jane.Address = &prototype.Address{StreetAddress: "451 2nd St.", City: "Miami", Country: "USA"}
		jane.Friends = []string{"Beth", "Sarah"}

		fmt.Printf("Name: %s, Address: %v, Friends: %v\n", john.Name, john.Address, john.Friends)
		fmt.Printf("Name: %s, Address: %v, Friends: %v\n", jane.Name, jane.Address, jane.Friends)
}


// prototypes require the ability to make deep copy of objects, types and structs
func betterPrototypeDeepCopyImpl() {
	john := prototype.Person{Name: "John",
		Address:     &prototype.Address{StreetAddress: "321 1st St.", City: "Houston", Country: "USA"},
		Friends: []string{"Chris", "Matt"},
	}

	// gives jane a copy of everything in john
	jane := john.BetterDeepCopy()
	jane.Name = "Jane"
	jane.Address = &prototype.Address{StreetAddress: "451 2nd St.", City: "Miami", Country: "USA"}
	jane.Friends = []string{"Beth", "Sarah"}

	fmt.Printf("Name: %s, Address: %v, Friends: %v\n", john.Name, john.Address, john.Friends)
	fmt.Printf("Name: %s, Address: %v, Friends: %v\n", jane.Name, jane.Address, jane.Friends)
}

func protoFactory() {
	john := prototype.NewMainOfficeEmployee("John", 100)
	jane := prototype.NewBranchOfficeEmployee("Jane", 300)

	fmt.Printf("Name: %s, Office: %v\n", john.Name, john.Office)
	fmt.Printf("Name: %s, Office: %v\n", jane.Name, jane.Office)
}

func adapterExample() {
	rc := adpater.NewRectangle(6, 4)
	a := adpater.VectorToRaster(rc)
	fmt.Println(adpater.DrawPoints(a))
}