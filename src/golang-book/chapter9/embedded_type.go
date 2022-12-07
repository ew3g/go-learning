package main

import "fmt"

type Person struct {
	Name string
}

func (p *Person) Talk() {
	fmt.Println("Hi, my name is", p.Name)
}

// Android -> An embedded type is like inheritance
// Here Android is also a Person
type Android struct {
	Person
	Model string
}

func main() {
	a := new(Android)
	a.Person.Name = "IDJK-34"
	a.Person.Talk()
}
