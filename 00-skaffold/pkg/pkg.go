package pkg

import "fmt"

type Nameable interface {
	Name() string
}

type Person struct {
	name string
}

func (p Person) Speak() {
	fmt.Printf("Hello from %v\n", p.name)
}

func NewPerson(name string) Person {
	return Person{
		name,
	}
}

func IsYourNameKumbi (nameable Nameable) bool {
	return nameable.Name() == "Kumbi"
}
