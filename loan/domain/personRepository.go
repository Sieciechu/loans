package domain

type PersonRepository interface {
	GetPersonByName(name string) (Person, error)
	AddPerson(p Person) error
	StorePerson(p Person) error
}