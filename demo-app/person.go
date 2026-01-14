package main

import "fmt"

// Person represents a person with name and age
type Person struct {
	Name string
	Age  int
}

// Introduce returns an introduction message for the person
func (p Person) Introduce() string {
	return fmt.Sprintf("Hi, I'm %s and I'm %d years old.", p.Name, p.Age)
}

// IsAdult checks if the person is 18 or older
func (p Person) IsAdult() bool {
	return p.Age >= 18
}

// HaveBirthday increments the person's age using a pointer receiver
func (p *Person) HaveBirthday() {
	p.Age++
}
