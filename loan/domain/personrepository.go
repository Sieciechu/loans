package domain

// PersonRepository - interface for getting and storing Person
type PersonRepository interface {
	GetPersonByName(name string) (Person, error)
	AddPerson(p Person) error
	StorePerson(p Person) error
}