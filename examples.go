package main



// basic struct
//p := new(person.Person)
//p.SetName("John Doe")
//p.SetAge(24)
//p.SetBirthday("01/02/1993")

//fmt.Println(p.GetName())
//fmt.Println(p.GetAge())
//fmt.Println(p.GetBirthday())

// struct implementing an interface
//human := person.CreateHuman(p)
//human.SetName("John Doe")
//human.SetAge(24)
//human.SetBirthday("01/02/1993")
//
//fmt.Println(human.GetName())
//fmt.Println(human.GetAge())
//fmt.Println(human.GetBirthday())

//hello := "hello"

// is a builder from the go SDK and it helps you concatenate strings into a buffer and then gets the
// final result
//sb := new(strings.Builder)
//sb.WriteString("<p>")
//sb.WriteString(hello)
//sb.WriteString("</p>")
//
//fmt.Println(sb.String())

// resets the content of the string builder(makes it empty and ready for reuse)
//words := []string{"Hello", "world", "test", "words", "builder"}
//sb.Reset()
//sb.WriteString("<ul>\n")
//for _, word := range words {
//	sb.WriteString("<li>")
//	sb.WriteString(word)
//	sb.WriteString("</li>\n")
//}
//sb.WriteString("</ul>")
//
//fmt.Println(sb.String())
//customHtmlBuilderExample()
//customPersonBuilderExample()
//functionalBuilder()
//customFunctionalFactory()
//customStructuredFactory()
//naivePrototypeDeepCopyImpl()
